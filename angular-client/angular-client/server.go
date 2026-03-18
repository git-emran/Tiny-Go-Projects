package main

import "net/http"

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
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
