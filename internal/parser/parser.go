package parser

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/opencars/opencars/internal/database"
	"github.com/opencars/opencars/pkg/model"
)

const (
	mappers   = 10
	reducers  = 10
	shufflers = 10 // Dont change. It closes channel.
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
	path = flag.String("path", "", "Path to CSV file")
)

func shuffler(input chan model.Operation, output chan []model.Operation, ready chan struct{}) {
	operations := make([]model.Operation, 0, batchSize)

	for {
		operation, ok := <-input
		if !ok {
			if len(operations) != 0 {
				output <- operations
				ready <- struct{}{}
			}
			break
		}

		if len(operations) < batchSize {
			operations = append(operations, operation)
			continue
		}

		output <- operations
		// TODO: Find anouther way to find too much memory allocation.
		operations = make([]model.Operation, 0, batchSize)
	}
}

func mapper(input chan []string, output chan model.Operation, ready chan struct{}) {
	for {
		msg, ok := <-input
		if !ok {
			ready <- struct{}{}
			break
		}

		output <- *model.NewOperation(msg)
	}
}

func reducer(input chan []model.Operation, output chan struct{}) {
	db := database.Must(database.DB())
	defer db.Close()

	for {
		operations, ok := <-input
		if !ok {
			output <- struct{}{}
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

	if *path == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(*path)
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

	mappersReady := make(chan struct{}, mappers)
	shufflersReady := make(chan struct{}, shufflers)
	reducersReady := make(chan struct{}, reducers)

	for i := 0; i < reducers; i++ {
		go reducer(batches, reducersReady)
	}

	for i := 0; i < shufflers; i++ {
		go shuffler(operations, batches, shufflersReady)
	}

	for i := 0; i < mappers; i++ {
		go mapper(rows, operations, mappersReady)
	}

	go mapperDispatcher(handler, rows)

	for i := 0; i < mappers; i++ {
		<-mappersReady
	}

	// Close channel.
	time.Sleep(time.Second)
	close(operations)

	// Wait for shufflers.
	for i := 0; i < shufflers; i++ {
		<-shufflersReady
	}

	// Close channel.
	time.Sleep(time.Second)
	close(batches)

	// Wait for reducers.
	for i := 0; i < reducers; i++ {
		<-reducersReady
	}

	log.Println("Execution time: ", time.Since(start))
}
