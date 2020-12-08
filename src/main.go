package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const appName = "2bitprogrammers/api_helloworld"
const appVersion = "2005.11a"
const appPort = "1234"

// RequestInfo contains all the data about the initial request
type RequestInfo struct {
	URI     string `json:"uri"`
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

// Response is what we will send back to the user
type Response struct {
	Date       time.Time   `json:"date"`
	StatusCode int         `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Data       string      `json:"data"`
	Errors     string      `json:"errors"`
	Request    RequestInfo `json:"request"`
}

// returnResponse - this will return a json response to the web client
func returnResponse(w http.ResponseWriter, method string, uri string, requestPayload string, status int, statusText string, data string, errors string) {
	sResponse := Response{Date: time.Now(), StatusCode: status, StatusText: statusText, Errors: errors}
	sResponse.Data = data
	sResponse.Request.URI = uri
	sResponse.Request.Method = method
	sResponse.Request.Payload = requestPayload

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if method == "OPTIONS" {
		log.Printf("%s %s %d %s", method, uri, http.StatusOK, requestPayload)
		return
	}

	log.Printf("%s %s %d %s", method, uri, status, requestPayload)
	if errors != "" {
		log.Printf("[ERROR] %s - %s", uri, errors)
	}

	joResponse, err := json.Marshal(sResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("[ERROR] Internal Server Error - Failed to parse Json.  %s", err), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		w.Write(joResponse)
		w.Write([]byte("\n\n\n"))
	} else {
		http.Error(w, string(joResponse), status)
	}

}

// handleStatusGet will handle all incoming status requests
func handleStatusGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI

	responsePayload := `{ "healthy": true}`
	returnResponse(w, method, uri, "", http.StatusOK, "OK", responsePayload, "")
}

// handleHelloWorldGet will handle all incoming /add requests
func handleHelloWorldGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI

	responsePayload := `"Hello World !!!"`
	returnResponse(w, method, uri, "", http.StatusOK, "OK", responsePayload, "")
}

func main() {
	fmt.Printf("%s v%s\n", appName, appVersion)
	fmt.Println("www.2BitProgrammers.com\nCopyright (C) 2020. All Rights Reserved.\n")
	log.Printf("Starting App on Port %s", appPort)

	http.HandleFunc("/status", handleStatusGet)
	http.HandleFunc("/hello", handleHelloWorldGet)
	http.HandleFunc("/hello-world", handleHelloWorldGet)
	http.HandleFunc("/hello_world", handleHelloWorldGet)
	http.HandleFunc("/helloworld", handleHelloWorldGet)
	http.HandleFunc("/helloWorld", handleHelloWorldGet)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
