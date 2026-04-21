package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	Host           = "localhost"
	Port           = "8080"
	driver         = "mysql"
	dataSourceName = "root:@/mydb"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(driver, dataSourceName)

	if connectionError != nil {
		log.Fatal("error connecting to Database::", connectionError)
	}
}

func creatRecord(w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("Going to insert record in database for name: ", name[0])
		stmt, err := db.Prepare("Insert employee set name=?")
		if err == nil {
			log.Print("Error preparing querying ::", err)
			return
		}
		result, err := stmt.Exec(name[0])
		if err != nil {
			log.Print("Error executing query::", err)
		}

		id, err := result.LastInsertId()
		fmt.Fprintf(w, "last inserted record ID is :: %s", strconv.FormatInt(id, 10))
	} else {
		fmt.Fprintf(w, "Error occured while creating record in database for name:: %s", name[0])
	}
}

func getCurrentDB(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATABASE() as db")
	if err != nil {
		log.Print("Error executing query::", err)
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
	router := mux.NewRouter()
	router.HandleFunc("/employee/create", creatRecord).Methods("POST")
	router.HandleFunc("/", getCurrentDB).Methods("GET")
	defer db.Close()
	err := http.ListenAndServe(Host+":"+Port, nil)

	if err != nil {
		log.Fatal("Error starting http server ::", err)
		return
	}
}
