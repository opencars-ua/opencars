package http

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/opencars/opencars/pkg/hsc"
	"github.com/opencars/opencars/pkg/model"
)

func VINHacker(uri string) {
	numOfErr := 0

	rand.Seed(time.Now().UTC().UnixNano())

	prefixes := [...]string{
		"CXI", "CAT", "CXX", "CXH", "CXM", "CXT", "AAC", "CAE", "AAE",
	}

	sdk := hsc.New(uri)

	for {
		prefix := prefixes[rand.Intn(len(prefixes))]
		digits := rand.Intn(1000000)
		code := fmt.Sprintf("%s%06d", prefix, digits)

		regs := make([]model.Registration, 0)
		err := Storage.Select(&regs, 1, "code = ?", code)
		if err != nil {
			log.Println(err.Error())
		}

		// TODO: Fix issue with duplicates in database!
		if len(regs) > 0 {
			continue
		}

		hscRegs, err := sdk.VehiclePassport(code)
		if err != nil {
			numOfErr++
			continue
		}

		for _, result := range hscRegs {
			var operations []model.Operation

			err := Storage.Select(&operations, 1, "number = ?", result.NRegNew)
			if err != nil {
				log.Printf("Select operations: %s\n", err)
				continue
			}

			// Save VIN code to the database.
			for i := range operations {
				operations[i].VIN = result.Vin
				if err := Storage.Update(&operations); err != nil {
					continue
				}
			}

			// Insert into database.
			obj := model.RegFromHSC(&result)
			if err := Storage.Insert(obj); err != nil {
				log.Printf("Insert registrations: %s\n", err)
				continue
			}

			log.Println(result.SDoc, result.NDoc, result.NRegNew, result.Vin)
		}
	}
}
