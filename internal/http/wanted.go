package http

import (
	"log"
	"net/http"

	"github.com/opencars/opencars/pkg/model"
	"github.com/opencars/opencars/pkg/translator"
)

type wantedPayload struct {
	Number string `schema:"number"`
	VIN    string `schema:"vin"`
	Limit  int    `schema:"limit"`
}

func (v *wantedPayload) Validate(r *http.Request) error {
	if v.Number == "" {
		return ErrInvalidNumber
	}

	if v.Limit < 1 {
		v.Limit = 1
	}

	return nil
}

func wanted(w http.ResponseWriter, r *http.Request) {
	payload := new(opsPayload)
	if err := decodeAndValidate(r, payload); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		log.Println(err)
		return
	}

	transport := make([]model.WantedTransport, 0)
	number := translator.ToUA(payload.Number)

	if err := Storage.Select(&transport, payload.Limit, "number = ?", number); err != nil {
		sendError(w, http.StatusInternalServerError, ErrInternal.Error())
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(transport); err != nil {
		sendError(w, http.StatusInternalServerError, ErrInternal.Error())
		log.Println(err)
		return
	}
}
