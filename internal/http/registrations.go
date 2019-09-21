package http

import (
	"log"
	"net/http"
	"regexp"

	"github.com/opencars/opencars/pkg/hsc"
	"github.com/opencars/opencars/pkg/model"
)

type regInfo struct {
	Code string `schema:"code,required"`
}

func (t regInfo) Validate(r *http.Request) error {
	match, _ := regexp.MatchString(`^\p{L}{3}\d{6}$`, t.Code)
	if !match {
		return ErrInvalidCode
	}

	return nil
}

type regsHandler struct{ *hsc.API }

func newRegsHandler(baseUrl string) *regsHandler {
	return &regsHandler{hsc.New(baseUrl)}
}

func (*regsHandler) save(regs []model.Registration) error {
	// Save registrations.
	if err := Storage.Insert(&regs); err != nil {
		return err
	}

	// Save VIN code for each Operation.
	for _, reg := range regs {
		var operation model.Operation

		err := Storage.Select(&operation, 1, "number = ?", reg.Number)
		if err != nil {
			return err
		}

		operation.VIN = reg.VIN

		if err := Storage.Update(&operation); err != nil {
			return err
		}
	}

	return nil
}

func (handler *regsHandler) fromDatabase(code string) ([]model.Registration, error) {
	var registrations []model.Registration

	err := Storage.Select(&registrations, 10, "code = ?", code)
	if err != nil {
		return nil, err
	}

	return registrations, nil
}

func (handler *regsHandler) fromRemote(code string) ([]model.Registration, error) {
	hscRegs, err := handler.VehiclePassport(code)
	if err != nil {
		return nil, err
	}

	regs := make([]model.Registration, len(hscRegs))
	for i, reg := range hscRegs {
		regs[i] = *model.RegFromHSC(&reg)
	}

	return regs, nil
}

func (handler *regsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload := new(regInfo)

	if err := decodeAndValidate(r, payload); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		log.Println(err)
		return
	}

	// TODO: We should handle this error.
	arr, _ := handler.fromDatabase(payload.Code)
	if len(arr) != 0 {
		if err := json.NewEncoder(w).Encode(arr); err != nil {
			sendError(w, http.StatusInternalServerError, ErrInternal.Error())
			log.Println(err)
		}
		return
	}

	arr, err := handler.fromRemote(payload.Code)
	if err != nil {
		sendError(w, http.StatusServiceUnavailable, ErrRemoteBroken.Error())
		log.Println(err)
		return
	}

	if len(arr) == 0 {
		sendError(w, http.StatusNotFound, ErrNotFound.Error())
		return
	}

	// Save.
	if err := handler.save(arr); err != nil {
		log.Printf("can not save registrations: %s\n", err)
	}

	if err := json.NewEncoder(w).Encode(arr); err != nil {
		sendError(w, http.StatusInternalServerError, ErrInternal.Error())
		log.Println(err)
		return
	}
}
