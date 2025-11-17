package main

import "github.com/gorilla/sessions"

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}
