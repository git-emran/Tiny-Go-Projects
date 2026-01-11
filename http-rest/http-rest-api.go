package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
		"getEmployee",
		"GET",
		"/employee/{id}",
		getEmployee,
	},

	Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	},

	Route{
		"updateEmployee",
		"PUT",
		"/employee/update",
		updateEmployee,
	},

	Route{
		"deleteEmployee",
		"DELETE",
		"/employee/delete",
		deleteEmployee,
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees []Employee
var employeesV1 []Employee
var employeesV2 []Employee

func init() {

	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Fizz", LastName: "Buzz"},
	}

	employeesV1 = Employees{
		Employee{Id: "1", FirstName: "Qux", LastName: "Klitt"},
		Employee{Id: "2", FirstName: "Qizz", LastName: "Blitt"},
	}

	employeesV2 = Employees{
		Employee{Id: "1", FirstName: "Beaux", LastName: "Neim"},
		Employee{Id: "2", FirstName: "Bill", LastName: "Veim"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(employees)
}

func getEmployeesV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeesV1)
}

func getEmployeesV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeesV2)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Print("error getting requested employee ::", err)
			}
		}
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Print("error occured while decoding employee data::", err)
		return
	}

	log.Printf("adding employee id:: %s with firstName as :: %s and lastName as:: %s", employee.Id, employee.FirstName, employee.LastName)

	employees = append(employees, Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName})
	json.NewEncoder(w).Encode(employees)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Print("Error occured while decoding employee data ::", err)
		return
	}

	var isUpsert = true

	for idx, emp := range employees {
		if emp.Id == employee.Id {
			isUpsert = false
			log.Printf("Updating employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			break
		}
	}

	if isUpsert {
		log.Printf("upserting employee id :: %s with firstName as :: %s and lastName as :: %s", employee.Id, employee.FirstName, employee.LastName)

		employees = append(employees, Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName})
	}

	json.NewEncoder(w).Encode(employees)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("Error occured while decoding employees data ::", err)
		return
	}
	log.Printf("deleting employee id :: %s with firstName as :: %s and lastName as :: %s", employee.Id, employee.FirstName, employee.LastName)

	index := GetIndex(employee.Id)
	employees = append(employees[:index], employees[index+1:]...)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func GetIndex(id string) int {
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			return i
		}
	}

	return -1
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)

	router.HandleFunc("/employees", getEmployees).Methods("GET")

	//v1
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/employees", getEmployeesV1).Methods("GET")

	//v2
	v2 := router.PathPrefix("/v2").Subrouter()
	v2.HandleFunc("/employees", getEmployeesV2).Methods("GET")

	log.Printf("Server running on http://%s:%s\n", CONN_HOST, CONN_PORT)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)

	if err != nil {
		log.Fatal("Error starting http server ::", err)
		return
	}
}
