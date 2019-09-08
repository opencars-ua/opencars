package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestOperations(t *testing.T) {
	uri := fmt.Sprintf("/vehicle/operations?number=АА9359РС")

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(operations)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "[]\n" {
		t.Errorf("got %s, expected %s", rr.Body.String(), "[]\n")
	}
}

func BenchmarkOperations(b *testing.B) {
	rand.Seed(time.Now().Unix())

	requests := make([]*http.Request, b.N)
	for i := 0; i < b.N; i++ {
		uri := fmt.Sprintf("/transport?number=BA%04dBT", rand.Intn(10000))
		req, _ := http.NewRequest("GET", uri, nil)
		requests[i] = req
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(operations)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(recorder, requests[i])
	}
}
