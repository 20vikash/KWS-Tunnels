package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	consts "tunnels/tunnels/consts/status"
	"tunnels/tunnels/models"

	"github.com/gorilla/websocket"
)

// Once login -> Redis(secret - UID)(with expiry)
// Create tunnel -> Populate postgres, Create nginx conf and reload(certbot if domain is custom)

var conns = make(map[string]*websocket.Conn) // Domain, web socket connection
var lock = sync.Mutex{}

// Tunnel channel
var tunChan chan models.Tunnel = make(chan models.Tunnel)

func (app *Application) WsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to parse form")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Before upgrading to WS.

	// 1. Get the secret from the header.
	secret := r.Header.Get("secret")

	// 2. Check the validity of the secret.
	uid, err := app.Store.InMemoryStore.GetUidFromSecret(r.Context(), secret)
	if err != nil {
		http.Error(w, "Invalid secret", http.StatusUnauthorized)
		return
	}

	// 3. Get the tunnel name from the body(check if the tunnel name matches with the secret).
	tunnel := r.Form.Get("tunnel")
	valid, err := app.Store.TunnelStore.ValidateTunnelFromUID(r.Context(), uid, tunnel)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !valid {
		log.Println("Cannot find given tunnel from the secret")
		http.Error(w, "Cannot find given tunnel from the secret", http.StatusBadRequest)
		return
	}

	// 4. Get the domain name from the tunnel name.
	domain, err := app.Store.TunnelStore.GetDomainFromTunnel(r.Context(), tunnel, uid)
	if err != nil {
		if err.Error() == consts.CANNOT_GET_DOMAIN {
			http.Error(w, "No domain", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Upgrade to WS and Populate the conns map once everything passed
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	app.AddConn(domain, conn)

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

func (app *Application) AddConn(domain string, ws *websocket.Conn) {
	lock.Lock()
	defer lock.Unlock()

	conns[domain] = ws
}

func (app *Application) RemoveConn(domain string) {
	lock.Lock()
	defer lock.Unlock()

	delete(conns, domain)
}
