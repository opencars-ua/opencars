package http

import (
	"log"
	"net/http"

	"github.com/opencars/opencars/pkg/model"
	"github.com/opencars/opencars/pkg/translator"
)

type opsPayload struct {
	Number string `schema:"number,required"`
	Limit  int    `schema:"limit"`
}

func (v *opsPayload) Validate(r *http.Request) error {
	if v.Number == "" {
		return ErrInvalidNumber
	}

	if v.Limit < 1 {
		v.Limit = 1
	}

	return nil
}

func operations(w http.ResponseWriter, r *http.Request) {
	payload := new(opsPayload)
	if err := decodeAndValidate(r, payload); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		log.Println(err)
		return
	}

	cars := make([]model.Operation, 0)
	number := translator.ToUA(payload.Number)

	if err := Storage.Select(&cars, payload.Limit, "number = ?", number); err != nil {
		sendError(w, http.StatusInternalServerError, ErrInternal.Error())
		log.Println(err)
		return
	}

	// TODO: Fetch VIN code from external services and save it.

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		sendError(w, http.StatusInternalServerError, ErrInternal.Error())
		log.Println(err)
		return
	}
}
