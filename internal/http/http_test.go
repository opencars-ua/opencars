package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Test struct{}

func (*Test) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	return nil
}

func (*Test) Update(
	model interface{},
) error {
	return nil
}

func (*Test) Insert(
	model interface{},
) error {
	return nil
}

func (*Test) Healthy() bool {
	return true
}

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("HEAD", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v, expected %v", status, http.StatusOK)
	}

	if rr.Body.String() != "" {
		t.Errorf("got %v, expected %v", rr.Body.String(), "")
	}
}

func TestMain(m *testing.M) {
	Storage = new(Test)
	code := m.Run()
	os.Exit(code)
}
