package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Fake struct{}

func (*Fake) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	return nil
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
	DB = &Fake{}
	handler := http.HandlerFunc(Handler)
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
	DB = &Fake{}
	handler := http.HandlerFunc(Handler)

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			b.Fail()
		}
	}
}
