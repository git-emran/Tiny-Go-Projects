package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	ID        string `json:"ID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{ID: "1", FirstName: "Uno", LastName: "Momento"},
		Employee{ID: "2", FirstName: "Dos", LastName: "Loco"},
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
	log.Printf("adding employee ID:: %s with first-name as :: %s and last-name as %s", employee.ID, employee.FirstName, employee.LastName)

	employees = append(employees, Employee{ID: employee.ID, FirstName: employee.FirstName, LastName: employee.LastName})
	json.NewEncoder(w).Encode(employees)
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	err := http.ListenAndServe(ConnHost+":"+ConnPort, router)

	if err != nil {
		log.Fatal("Error starting http server ::", err)
		return
	}

}
