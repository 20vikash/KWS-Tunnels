package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Once login -> Redis(secret - UID)(with expiry)
// Create tunnel -> Populate postgres, Create nginx conf and reload(certbot if domain is custom)
// Run tunnel -> (client sends secret, tunnel name) -> get the domain from the tunnel name
// Destroy tunnel -> Destroy everything

// Conns needed for tunnel server
// 1. Postgres (To get domain name from tunnel name)

var conns = make(map[string]*websocket.Conn) // Domain, web socket connection
var lock = sync.Mutex{}

func (app *Application) WsHandler(w http.ResponseWriter, r *http.Request) {
	// Before upgrading to WS.
	// 1. Get the secret from the header.
	// 2. Check the validity of the secret.
	// 3. Get the tunnel name from the body.
	// 4. Get the domain name from the tunnel name.
	// 5. Populate the conns map once everything passed

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	// Listen for incoming messages
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received: %s\\n", message)
		// Echo the message back to the client
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
