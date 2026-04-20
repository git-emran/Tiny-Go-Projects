package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	http.HandleFunc("/", getCurrentDB)
	defer db.Close()
	err := http.ListenAndServe(Host+":"+Port, nil)

	if err != nil {
		log.Fatal("Error starting http server ::", err)
		return
	}
}
