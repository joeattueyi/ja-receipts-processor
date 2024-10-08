package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var idToPoints = make(map[uuid.UUID]int)

type Application struct {
	port   int
	router *mux.Router
}

func main() {

	app := Application{port: 8080}
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handleProcess).
		Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlePoints).
		Methods("GET")
	app.router = r
	http.Handle("/", r)
	log.Println("Starting on port: ", app.port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", app.port), nil))
}

func handleProcess(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("The receipt is invalid"))
		return
	}

	receipt := Receipt{}

	json.Unmarshal(body, &receipt)
	validator := NewValidator()
	ValidateReceipt(validator, &receipt)
	if !validator.Valid() {
		e := map[string]interface{}{
			"message": "The receipt is invalid",
			"detail":  validator.Errors,
		}
		js, _ := json.MarshalIndent(e, " ", "\t")
		js = append(js, '\n')

		w.WriteHeader(400)
		w.Write(js)
		return
	}

	uuid := uuid.NewSHA1(uuid.Max, body)

	response := struct {
		Id string `json:"id"`
	}{
		uuid.String(),
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("UUID generated could not be marshaled"))
		return
	}

	if _, exists := idToPoints[uuid]; !exists {
		idToPoints[uuid] = receipt.computePoints()
	}

	w.Write(json)

}

func handlePoints(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	uuid, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("No receipt found for that id"))
		return
	}

	points, ok := idToPoints[uuid]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("No receipt found for that id"))
		return
	}

	response := struct {
		Points int64 `json:"points"`
	}{
		int64(points),
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("No receipt found for that id"))
		return
	}

	w.Write(json)
}
