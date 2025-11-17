package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func home(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintln(w, "you have successfully logged in")

}

func logout(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

}
