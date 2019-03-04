package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/thomasraydeniscool/tinychip"
)

func TestGetOrders(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	GetOrders(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}
