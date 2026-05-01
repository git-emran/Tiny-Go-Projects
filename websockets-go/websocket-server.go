package main

import "github.com/gorilla/websocket"

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

type Message struct {
	Message string `json:"message"`
}
