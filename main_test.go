package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleRoutes(t *testing.T) {
	app := Application{port: 8080}
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handleProcess).
		Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlePoints).
		Methods("GET")
	app.router = r

	ts := httptest.NewServer(app.router)
	defer ts.Close()

	resp, err := ts.Client().Post(ts.URL+"/receipts/process", "application/json", strings.NewReader(
		`{
			"retailer": "Target",
			"purchaseDate": "2022-01-01",
			"purchaseTime": "13:01",
			"items": [
			  {
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
			  },{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
			  },{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			  },{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			  },{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
			  }
			],
			"total": "35.35"
		  }`))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Cannot read response body")
	}

	payload := struct{ Id string }{}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Fatal("unable to parse receipt")
	}

	resp, err = ts.Client().Get(ts.URL + fmt.Sprintf("/receipts/%s/points", payload.Id))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Cannot read response body")
	}
	points := struct{ Id string }{}
	err = json.Unmarshal(body, &points)
	if err != nil {
		t.Fatal("Unable to parse")
	}

}

func TestFailure(t *testing.T) {
	app := Application{port: 8080}
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handleProcess).
		Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlePoints).
		Methods("GET")
	app.router = r

	ts := httptest.NewServer(app.router)
	defer ts.Close()

	resp, err := ts.Client().Post(ts.URL+"/receipts/process", "application/json", strings.NewReader(
		``))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 400 {
		t.Errorf("got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Cannot read response body")
	}

	payload := struct{ Id string }{}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Fatal("unable to parse receipt")
	}

	resp, err = ts.Client().Get(ts.URL + fmt.Sprintf("/receipts/%s/points", payload.Id))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 404 {
		t.Errorf("got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
}
