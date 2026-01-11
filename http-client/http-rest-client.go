package main

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8090"
)

const WEB_SERVICE_HOST string = "http://localhost:8080"

type employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
