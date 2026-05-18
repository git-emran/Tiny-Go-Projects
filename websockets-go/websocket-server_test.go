package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"golang.org/x/tools/playground/socket"
)

func TestWebSocketServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleClients()))
	defer server.Close()
	u := "ws" + strings.TrimPrefix(server.URL, "http")
	socket, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer socket.Close()
	m := Message{Message: "hello"}
}
