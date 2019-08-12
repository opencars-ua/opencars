package parser

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-pg/pg"
	"github.com/opencars/opencars/internal/config"
	"github.com/opencars/opencars/internal/storage"

	"github.com/opencars/opencars/pkg/model"
)

const (
	mappers   = 10
	reducers  = 10
	shufflers = 10 // Don't change. It closes channel.
	batchSize = 5000
)

type HandlerCSV struct {
	reader *csv.Reader
}

func (h *HandlerCSV) ReadN(amount int) ([][]string, error) {
	result := make([][]string, 0)

	for i := 0; i < amount; i++ {
		record, err := h.reader.Read()
		if err == io.EOF {
			return result, err
		}

		if err != nil {
			return nil, err
		}

		result = append(result, record)
	}

	return result, nil
}

var (
	sourcePath = flag.String("csv", "", "Path to CSV file")
	configPath = flag.String("config", "./config/opencars.toml", "Path to configuration file")
)

func shuffler(wg *sync.WaitGroup, input chan model.Operation, output chan []model.Operation) {
	operations := make([]model.Operation, 0, batchSize)

	for {
		operation, ok := <-input
		if !ok {
			if len(operations) != 0 {
				output <- operations
			}
			wg.Done()
			break
		}

		if len(operations) < batchSize {
			operations = append(operations, operation)
			continue
		}

		output <- operations

		// TODO: Find another way to find too much memory allocation.
		operations = make([]model.Operation, 0, batchSize)
	}
}

func mapper(wg *sync.WaitGroup, input chan []string, output chan model.Operation) {
	for {
		msg, ok := <-input
		if !ok {
			wg.Done()
			break
		}

		output <- *model.NewOperation(msg)
	}
}

func reducer(wg *sync.WaitGroup, db *pg.DB, input chan []model.Operation) {
	for {
		operations, ok := <-input
		if !ok {
			wg.Done()
			break
		}

		if err := db.Insert(&operations); err != nil {
			log.Fatal(err)
		}

		log.Printf("Done: %d\n", len(operations))
	}
}

func mapperDispatcher(handler HandlerCSV, output chan []string) {
	for {
		msgs, err := handler.ReadN(5000)

		if err == nil || err == io.EOF {
			for _, msg := range msgs {
				output <- msg
			}
		}

		if err == io.EOF {
			close(output)
			break
		}

		if err != nil {
			log.Println(err)
			close(output)
			break
		}
	}
}

func Run() {
	flag.Parse()
	start := time.Now()

	conf, err := config.New(*configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := storage.New(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = storage.Migrate(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *sourcePath == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(*sourcePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Skip header line.
	if _, err := csvReader.Read(); err != nil {
		log.Panic(err.Error())
	}

	handler := HandlerCSV{reader: csvReader}
	rows := make(chan []string, 100000)
	operations := make(chan model.Operation, 100000)
	batches := make(chan []model.Operation, 10000)

	mapperWg := sync.WaitGroup{}
	shufflersWg := sync.WaitGroup{}
	reducersWg := sync.WaitGroup{}

	for i := 0; i < reducers; i++ {
		reducersWg.Add(1)
		go reducer(&reducersWg, db, batches)
	}

	for i := 0; i < shufflers; i++ {
		shufflersWg.Add(1)
		go shuffler(&shufflersWg, operations, batches)
	}

	for i := 0; i < mappers; i++ {
		mapperWg.Add(1)
		go mapper(&mapperWg, rows, operations)
	}

	go mapperDispatcher(handler, rows)

	mapperWg.Wait()

	// Close channel.
	time.Sleep(time.Second)
	close(operations)

	shufflersWg.Wait()

	// Close channel.
	time.Sleep(time.Second)
	close(batches)

	time.Sleep(time.Second)
	// Wait for reducers.
	reducersWg.Wait()

	log.Println("Execution time: ", time.Since(start))
}
