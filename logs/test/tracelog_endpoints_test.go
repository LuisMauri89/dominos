package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"dominos.com/logs"
)

func TestEmptyTlogs(t *testing.T) {
	ClearTableTlogs()

	req, _ := http.NewRequest("GET", "/tlogs/", nil)
	response := ExecuteRequest(req)

	actual := response.Code
	want := http.StatusOK

	if want != actual {
		t.Errorf("Expected response code %d. Got %d\n", want, actual)
	}

	var res logs.FindAllResponse
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		t.Errorf("Can not parse json response body. Body: %s", response.Body.String())
	}

	if count := len(res.Tlogs); count != 0 {
		t.Errorf("Expected an empty array. Got len %d", count)
	}
}

func TestOneTlog(t *testing.T) {
	SeedTableTlogsOneTlog()

	req, _ := http.NewRequest("GET", "/tlogs/", nil)
	response := ExecuteRequest(req)

	actual := response.Code
	want := http.StatusOK

	if want != actual {
		t.Errorf("Expected response code %d. Got %d\n", want, actual)
	}

	var res logs.FindAllResponse
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		t.Errorf("Can not parse json response body. Body: %s", response.Body.String())
	}

	if count := len(res.Tlogs); count != 1 {
		t.Errorf("Expected an array with one element. Got len %d", count)
	}
}

func TestCreateTlog(t *testing.T) {
	ClearTableTlogs()

	var jsonStr = []byte(`{"tlog": {"serviceName":"DELIVERIES", "caller": "Delivered Method", "event": "DELIVER", "extra": "to Luis Mauri."}}`)
	req, _ := http.NewRequest("POST", "/tlogs/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := ExecuteRequest(req)

	actual := response.Code
	want := http.StatusOK

	if want != actual {
		t.Errorf("Expected response code %d. Got %d\n", want, actual)
	}
}
