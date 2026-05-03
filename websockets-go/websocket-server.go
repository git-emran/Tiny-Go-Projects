package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type Message struct {
	Message string `json:"message"`
}

func handleClients(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade error ::", err)
		return
	}

	clients[conn] = true

	go func(c *websocket.Conn) {
		defer func() {
			delete(clients, c)
			c.Close()
		}()

		for {
			var msg Message
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Panicln("read error ::", err)
				return
			}
			broadcast <- msg
		}
	}(conn)
}

func broadcastMessagesToClients() {
	for {
		message := <-broadcast

		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("Error occured while writing message to client  %v ::", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	go broadcastMessagesToClients()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/echo", handleClients)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
