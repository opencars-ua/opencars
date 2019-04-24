package parser

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"

	"github.com/opencars/opencars/internal/database"
	"github.com/opencars/opencars/pkg/model"
)

type HandlerCSV struct {
	db     *pg.DB
	reader *csv.Reader
}

func (h *HandlerCSV) ReadN(amount int) ([]model.Operation, error) {
	result := make([]model.Operation, 0)

	for i := 0; i < amount; i++ {
		record, err := h.reader.Read()
		if err == io.EOF {
			return result, err
		}

		if err != nil {
			return nil, err
		}

		car := model.NewOperation(record)
		if car.Valid() {
			result = append(result, *car)
		}
	}

	return result, nil
}

var (
	path = flag.String("path", "", "Path to XSV file")
)

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
	createdCars := 0

	db := database.Must(database.DB())
	defer db.Close()

	// Skip header line.
	if _, err := csvReader.Read(); err != nil {
		log.Println(err.Error())
	}

	handler := HandlerCSV{
		db:     db,
		reader: csvReader,
	}

	N := 10000

	for {
		cars, readErr := handler.ReadN(N)

		err = db.Insert(&cars)
		if err != nil {
			log.Println(err)
		} else {
			createdCars += N
			log.Println(createdCars)
		}

		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("Execution time: ", time.Since(start))
}
