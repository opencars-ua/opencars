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
	Mappers = 1
	Reducers = 1
	Shufflers = 1
	BatchSize = 20000
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
	path = flag.String("path", "", "Path to XSV file")
)

func shuffler(input chan model.Operation, output chan []model.Operation) {
	operations := make([]model.Operation, 0, BatchSize)

	for {
		operation, open := <- input
		if !open {
			if len(operations) != 0 {
				output <- operations
			}

			return
		}

		if len(operations) < BatchSize {
			operations = append(operations, operation)
			continue
		}

		output <- operations
		operations = operations[:0]
	}
}

func mapper(input chan []string, output chan model.Operation) {
	for {
		msg, opened := <- input
		if !opened {
			return
		}

		output <- *model.NewOperation(msg)
	}
}

func reducer(input chan []model.Operation, output chan struct{}) {
	db := database.Must(database.DB())
	defer db.Close()

	for {
		operations, open := <-input
		if !open {
			output <- struct{}{}
			return
		}

		if err := db.Insert(&operations); err != nil {
			log.Println(len(operations))
			log.Println(err)
		}
	}
}

func mapperDispatcher(handler HandlerCSV, output chan []string, red chan model.Operation) {
	for {
		msgs, err := handler.ReadN(10000)
		if err == nil || err == io.EOF {
			for _, msg := range msgs {
				output <- msg
			}
		}

		if err != nil {
			close(output)
			time.Sleep(time.Second)
			close(red)
			return
		}
	}
}

func Run() {
	start := time.Now()

	flag.Parse()

	if *path == "" {
		panic("empty path")
	}

	file, err := os.Open(*path)
	if err != nil {
		panic(err.Error())
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// Skip header line.
	if _, err := csvReader.Read(); err != nil {
		log.Println(err.Error())
	}

	handler := HandlerCSV{reader: csvReader}
	rows := make(chan []string, 100000)
	operations := make(chan model.Operation, 100000)
	batches := make(chan []model.Operation, 100)
	ready := make(chan struct{}, Reducers)

	for i := 0; i < Reducers; i++ {
		go reducer(batches, ready)
	}

	for i := 0; i < Shufflers; i++ {
		go shuffler(operations, batches)
	}

	for i := 0; i < Mappers; i++ {
		go mapper(rows, operations)
	}

	go mapperDispatcher(handler, rows, operations)

	for i := 0; i < Reducers; i++ {
		<- ready
	}

	fmt.Println("Execution time: ", time.Since(start))
}
