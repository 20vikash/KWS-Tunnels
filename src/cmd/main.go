package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Application struct {
	Host string
	Port string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Check for the session information before authorizing
		return true
	},
}

func main() {
	app := &Application{
		Host: "0.0.0.0",
		Port: "8081",
	}

	fmt.Println("WebSocket server started on :8081")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", app.Host, app.Port), NewRouter(app))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
