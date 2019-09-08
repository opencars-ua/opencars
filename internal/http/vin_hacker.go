package http

import (
	"fmt"
	"log"
	"sync"

	"github.com/opencars/opencars/pkg/hsc"
	"github.com/opencars/opencars/pkg/model"
)

const (
	sizeOfBulkInsert = 1000
	sizeOfBulkFetch  = 100
	total            = 100000
)

// Event helps decouple logic using golang channels.
type Event struct {
	Registrations []model.Registration
	Processed     uint32
}

// NewEvent returns new instance of Event.
func NewEvent(
	registrations []model.Registration,
	processed uint32,
) *Event {
	return &Event{
		Registrations: registrations,
		Processed:     processed,
	}
}

func worker(
	wg *sync.WaitGroup,
	from, to uint32,
	events chan *Event,
	prefix, uri string,
) {
	sdk := hsc.New(uri)

	regs := make([]model.Registration, 0)
	for i := from; i < to; i++ {
		if len(regs) >= sizeOfBulkFetch {
			events <- NewEvent(regs, uint32(len(regs)))
			regs = make([]model.Registration, 0)
		}

		code := fmt.Sprintf("%s%06d", prefix, i)
		hscRegs, err := sdk.VehiclePassport(code)
		if err != nil {
			log.Printf("code: %s. Remote error: %s\n", code, err)
			continue
		}

		for _, hscReg := range hscRegs {
			obj := model.RegFromHSC(&hscReg)
			regs = append(regs, *obj)
		}
	}

	if len(regs) != 0 {
		events <- NewEvent(regs, uint32(len(regs)))
	}

	wg.Done()
}

func batch(events chan *Event) {
	// Insert collects 50K and then insert all of them at once.
	collected := make([]model.Registration, 0, sizeOfBulkInsert*2)
	inserted, processed := 0, uint32(0)

	for {
		event, ok := <-events
		processed += event.Processed

		collected = append(collected, event.Registrations...)
		if len(collected) < sizeOfBulkInsert && ok {
			continue
		}

		if err := Storage.Insert(&collected); err != nil {
			log.Fatal(err)
		}
		// Increment number of inserted rows.
		inserted += len(collected)
		// Clean up.
		collected = collected[:0]

		log.Printf("Processed: %d/%d\n", processed, total)
		log.Printf("Inserted: %d\n", inserted)

		if !ok {
			break
		}
	}
}

// VINHacker starts process of VIN codes extraction from third-party service.
func VINHacker(prefix, uri string, threads uint16) {
	events := make(chan *Event, threads*2)
	numPerThread := 1000000 / uint32(threads)
	wg := sync.WaitGroup{}

	for i := uint32(0); i < uint32(threads); i++ {
		from := i * numPerThread
		to := from + numPerThread
		wg.Add(1)
		go worker(&wg, from, to, events, prefix, uri)
	}

	go batch(events)
	go batch(events)

	wg.Wait()
	close(events)
}
