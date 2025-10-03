package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"tunnels/tunnels/models"

	"github.com/google/uuid"
)

func (app *Application) HandleTunnelRequest(w http.ResponseWriter, r *http.Request) {
	// Create a new uuid
	uuid := uuid.NewString()

	// Dump http message into byte array
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Failed to dump request")
		http.Error(w, "Failed to dump request", http.StatusInternalServerError)
		return
	}

	tunnel := models.Tunnel{
		RequestID: uuid,
		Data:      dump,
	}

	// Get the host header
	host := r.Host

	// Send the message to the corresponding ws connection.
	ws := conns[host]
	ws.WriteJSON(tunnel)

	//TODO: Wait on the channel for response
}
