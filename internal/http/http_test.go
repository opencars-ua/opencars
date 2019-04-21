package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func (*Test) Healthy() bool {
	return true
}

func TestTransport(t *testing.T) {
	Storage = new(Test)
	uri := "localhost:8080"

	t.Run("success", func(t *testing.T) {
		params := url.Values{}
		params.Add("number", "BA2927BT")
		assert.HTTPSuccess(t, Transport, "GET", uri, params)
	})

	t.Run("missing number", func(t *testing.T) {
		assert.HTTPError(t, Transport, "GET", uri, nil)
	})
}

func BenchmarkHandler(b *testing.B) {
	rand.Seed(time.Now().Unix())
	requests := make([]*http.Request, b.N)

	for i := 0; i < b.N; i++ {
		number := 1000 + rand.Uint64()%8999
		uri := fmt.Sprintf("/transport?number=BA%dBT", number)
		req, _ := http.NewRequest(
			"GET",
			uri,
			nil,
		)
		requests[i] = req
	}

	recorder := httptest.NewRecorder()
	Storage = new(Test)
	handler := http.HandlerFunc(Transport)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(recorder, requests[i])
	}
}
