package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ConnHost = "localhost"
	ConnPort = "8080"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},

	Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Uno", LastName: "Momento"},
		Employee{Id: "2", FirstName: "Dos", LastName: "Loco"},
	}

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)

}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("Error occured while decoing employee data::", err)
		return
	}
	log.Printf("adding employee id:: %s with first-name as :: %s and last-name as %s", employee.Id, employee.FirstName, employee.LastName)

}
