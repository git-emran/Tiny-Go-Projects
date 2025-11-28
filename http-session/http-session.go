package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/boj/redistore.v1"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var store *redistore.RediStore

func init() {
	var err error
	store, err = redistore.NewRediStore(10, "tcp", "localhost:6379", "", []byte("secret-key"))

	if err != nil {
		log.Fatal("error connecting to redis store: ", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Unauthorized to view the page", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "Welcome Home")

}

func login(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	if err := session.Save(r, w); err != nil {
		log.Fatalf("Error saving sessions: %v", err)
	}

	fmt.Fprintln(w, "You have successfully logged in")

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintln(w, "you have successfully logged out")

}

func main() {
	defer store.Close()
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	log.Println("Server started at http://" + CONN_HOST + ":" + CONN_PORT)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}

}
