package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	Host           = "localhost"
	Port           = "8080"
	driver         = "mysql"
	dataSourceName = "root:@/mydb"
)

type Employee struct {
	ID   int    `json:"uid"`
	Name string `json:"name"`
}

var (
	db              *sql.DB
	connectionError error
)

func init() {
	db, connectionError = sql.Open(driver, dataSourceName)

	if connectionError != nil {
		log.Fatal("error connecting to Database::", connectionError)
	}
}

func readRecords(w http.ResponseWriter, r *http.Request) {
	log.Print("reading database records")
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Print("Error executing select query ::", err)
		return
	}

	employees := []Employee{}
	for rows.Next() {
		var uid int
		var name string
		err = rows.Scan(&uid, &name)
		if err != nil {
			log.Print("Error scanning the name")
		}

		employee := Employee{ID: uid, Name: name}
		employees = append(employees, employee)
	}
	json.NewEncoder(w).Encode(employees)
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]

	if !ok || len(name) == 0 {
		log.Print("Missing 'name' in parameter request ")
		http.Error(w, "Query parameter name is required", http.StatusBadRequest)
		return
	}
	log.Print("Going to insert record in database for name: ", name[0])
	result, err := db.Exec("INSERT INTO employee (name) VALUES (?)", name[0])
	if err != nil {
		log.Printf("Error executing query %v", err)
		return
	}
	id, _ := result.LastInsertId()

	fmt.Fprintf(w, "Last insert record ID is :: %d", id)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	// store the request and get the id
	vars := mux.Vars(r)
	id := vars["id"]
	// store the query in 'vals' and print the name
	vals := r.URL.Query()
	name, ok := vals["name"]

	if ok {
		log.Print("Updating record in database for id::", id)
		// Making the db query
		stmt, err := db.Prepare("UPDATE employee SET name=? where id=?")
		if err != nil {
			log.Print("Error updating employee name ::", err)
			return
		}
		result, err := stmt.Exec(name[0], id)
		if err != nil {
			log.Print("Error while executing query ::", err)
		}
		rowsAffected, err := result.RowsAffected()
		fmt.Printf("Number of rows updated in database %d", rowsAffected)
		if err != nil {
			log.Print("Error printing result")
		}
	} else {
		fmt.Printf("Error while updating record in database for id :: %s", id)
	}
}

func getCurrentDB(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATABASE() as db")
	if err != nil {
		log.Printf("Error executing query:: %v", err)
		return
	}
	defer rows.Close()

	var db string
	for rows.Next() {
		rows.Scan(&db)
	}

	fmt.Fprintf(w, "Current database is:: %s", db)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employee/update/{id}", updateRecord).Methods("PUT")
	router.HandleFunc("/employee/create", createRecord).Methods("POST")
	router.HandleFunc("/employees", readRecords).Methods("GET")
	router.HandleFunc("/", getCurrentDB).Methods("GET")
	defer db.Close()
	err := http.ListenAndServe(Host+":"+Port, router)
	if err != nil {
		log.Fatal("Error starting http server ::", err)
		return
	}
}
