package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	ID   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{ID: "1", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occured while executing the template %v", err)
		return
	}
}

func main() {

	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server", err)
		return
	}
}
