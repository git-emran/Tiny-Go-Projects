package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/resty.v1"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8090"
)

const WEB_SERVICE_HOST string = "http://localhost:8080"

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	response, err := resty.R().Get(WEB_SERVICE_HOST + "/employee")
	if err != nil {
		log.Printf("Error getting data from the webservice %s::", err)
		return
	}

	print(response, err)
	fmt.Fprint(w, response.String())

}
