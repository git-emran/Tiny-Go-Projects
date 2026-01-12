package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

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

func getEmployees(w http.ResponseWriter, r *http.Request) {
	response, err := resty.R().Get(WEB_SERVICE_HOST + "/employees")
	if err != nil {
		log.Printf("Error getting data from the webservice %s::", err)
		return
	}

	print(response, err)
	fmt.Fprint(w, response.String())

}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	decodingErr := json.NewDecoder(r.Body).Decode(&employee)

	if decodingErr != nil {
		log.Print("Error occured while decoidng employee", decodingErr)
		return
	}

	log.Print("adding employee id ::", employee.Id, employee.FirstName, employee.LastName)
	response, err := resty.R().SetHeader("Content-Type", "application/json").SetBody(Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName}).Post(WEB_SERVICE_HOST + "/employee/add")
	if err != nil {
		log.Print("error occured while adding employee ::", err)
		return
	}

	print(response, err)
	fmt.Print(w, response.String())

}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	employee := router.PathPrefix("/employee").Subrouter()
	employee.HandleFunc("/add", addEmployee).Methods("POST")
}
