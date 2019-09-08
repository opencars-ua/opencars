package http

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/opencars/opencars/pkg/model"
)

var (
	registrationsPath = "../../test/registrations.json"
)

func TestRegsHandler_ServeHTTP(t *testing.T) {
	uri := fmt.Sprintf("/vehicle/registrations?code=CXI200455")

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	payload, err := ioutil.ReadFile(registrationsPath)
	if err != nil {
		t.Fatal()
	}

	rr := httptest.NewRecorder()

	t.Run("returns registrations", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if _, err := w.Write(payload); err != nil {
					t.Fatal()
				}
			}),
		)
		defer server.Close()

		handler := newRegsHandler(server.URL)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("got %v, expected %v", status, http.StatusOK)
		}

		expected := make([]model.Registration, 0)
		if err := json.NewEncoder(rr.Body).Encode(&expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("remote server is not available", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "", http.StatusServiceUnavailable)
			}),
		)
		defer server.Close()

		handler := newRegsHandler(server.URL)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusServiceUnavailable {
			t.Errorf("got %v, expected %v", status, http.StatusServiceUnavailable)
		}

		expected := "{\"error\":\"remote server is not available\"}\n"
		if rr.Body.String() != expected {
			t.Errorf("got %v, expected %v", rr.Body.String(), expected)
		}
	})
}

func BenchmarkRegsHandler_ServeHTTP(b *testing.B) {
	rand.Seed(time.Now().Unix())

	requests := make([]*http.Request, b.N)
	for i := 0; i < b.N; i++ {
		uri := fmt.Sprintf("/vehicle/registrations?code=CXI%06d", rand.Intn(1000000))
		req, _ := http.NewRequest("GET", uri, nil)
		requests[i] = req
	}

	payload, err := ioutil.ReadFile(registrationsPath)
	if err != nil {
		b.Fatal()
	}

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write(payload); err != nil {
				b.Fatal()
			}
		}),
	)
	defer server.Close()

	rr := httptest.NewRecorder()
	handler := newRegsHandler(server.URL)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, requests[i])
	}
}
