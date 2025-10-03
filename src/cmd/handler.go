package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
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

	// Tunnel channel
	var tunChan chan models.Tunnel = make(chan models.Tunnel)
	tunChans[uuid] = tunChan
	defer delete(tunChans, uuid)

	// Send the message to the corresponding ws connection.
	ws := conns[host]
	ws.Write.Lock()
	ws.Conn.WriteJSON(tunnel)
	ws.Write.Unlock()

	// Wait on the channel for response
	select {
	case response := <-tunChan:
		w.Write(response.Data)
	case <-time.After(10 * time.Second): // Timeout from upstream
		http.Error(w, "Upstream timeout", http.StatusGatewayTimeout)
		return
	}
}
