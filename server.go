package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// define a struct to represent the data
type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// handler function for GET request. Respond with the current day of the week
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	//create data
	data := Data{
		Key:   "Day of week",
		Value: time.Now().Weekday().String(),
	}
	// convert to JSON format
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set header and write response
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

// handler function for post request. send back what was received
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// parse request to Data struct
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Received data: %+v", data)

}

func main() {

	http.HandleFunc("/get", handleGetRequest)
	http.HandleFunc("/post", handlePostRequest)

	fmt.Println("Server Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
