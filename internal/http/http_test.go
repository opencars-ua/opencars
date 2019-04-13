package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestAdapter struct{}

func (*TestAdapter) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	return nil
}

func (*TestAdapter) Healthy() bool {
	return true
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(
		"GET",
		"/transport?number=BA2927BT",
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	DB = &TestAdapter{}
	handler := http.HandlerFunc(Transport)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fail()
	}

	// Check the response body is what we expect.
	expected := "[]\n"
	if rr.Body.String() != expected {
		t.Fail()
	}
}

func BenchmarkHandler(b *testing.B) {
	req, err := http.NewRequest(
		"GET",
		"/transport?number=BA2927BT",
		nil,
	)

	if err != nil {
		b.Fatal(err)
	}

	rr := httptest.NewRecorder()
	DB = &TestAdapter{}
	handler := http.HandlerFunc(Transport)

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			b.Fail()
		}
	}
}
