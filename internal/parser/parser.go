package parser

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-pg/pg"
	"github.com/shal/opencars/internal/database"
	"github.com/shal/opencars/pkg/models"

)

type HandlerCSV struct {
	db     *pg.DB
	reader *csv.Reader
}

func (h *HandlerCSV) ReadN(amount int) ([]models.Transport, error) {
	result := make([]models.Transport, amount)

	for i := 0; i < amount; i++ {
		record, err := h.reader.Read()
		if err == io.EOF {
			return result, io.EOF
		}

		if err != nil {
			fmt.Println("Something went wrong, while reading!")
			continue
		}

		car := models.NewTransportFromCSV(record)
		if !car.Valid() {
			continue
		}

		result[i] = *car
	}

	return result, nil
}

func Run() {
	start := time.Now()

	path := flag.String("path", "", "Path to xsv file")
	flag.Parse()

	if strings.TrimSpace(*path) == "" {
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
	csvReader.Read()

	handler := HandlerCSV{
		db:     db,
		reader: csvReader,
	}

	N := 10000

	for {
		cars, readErr := handler.ReadN(N)

		err = db.Insert(&cars)
		if err != nil {
			fmt.Println(err)
		} else {
			createdCars += N
			fmt.Println(createdCars)
		}

		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("Execution time: ", time.Since(start))
}
