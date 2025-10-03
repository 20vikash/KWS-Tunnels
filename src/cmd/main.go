package main

import (
	"fmt"
	"net/http"
	env "tunnels/tunnels/internal"
	"tunnels/tunnels/internal/database"
	"tunnels/tunnels/internal/store"

	"github.com/gorilla/websocket"
)

type Application struct {
	Host  string
	Port  string
	Store *store.Storage
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Check for the session information before authorizing
		return true
	},
}

func main() {
	// Create database connections.
	pg := database.Pg{
		User:     env.GetDBUserName(),
		Password: env.GetDBPassword(),
		Host:     env.GetDBHost(),
		Port:     env.GetDBPort(),
		Name:     env.GetDBName(),
	}
	pgCon := pg.GetNewDBConnection()

	redis := database.RedisDB{
		Addr:     fmt.Sprintf("%s:%s", env.GetRedisHost(), env.GetRedisPort()),
		Password: env.GetRedisPassword(),
		DB:       2,
	}
	redisCon := redis.Connect()

	store := store.NewStore(pgCon, redisCon)

	app := &Application{
		Host:  "0.0.0.0",
		Port:  "8081",
		Store: store,
	}

	fmt.Println("WebSocket server started on :8081")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", app.Host, app.Port), NewRouter(app))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
