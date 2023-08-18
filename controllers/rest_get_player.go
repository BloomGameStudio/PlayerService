package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"nhooyr.io/websocket"
)

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
    playerServiceURL := "ws://localhost:1323/ws/player"
    fmt.Printf("Got a request lets return some data.\n")
    conn, _, err := websocket.Dial(r.Context(), playerServiceURL, nil)
    if err != nil {
        http.Error(w, "Error connecting to WebSocket", http.StatusInternalServerError)
        return
    }
    defer conn.Close(websocket.StatusNormalClosure, "")
	_, msg, err := conn.Read(r.Context())
	if err != nil {
		log.Printf("Error reading message: %v", err)
		http.Error(w, "Error reading message from WebSocket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(msg)
	if err != nil {
		log.Printf("Error forwarding response data: %v", err)
	}
}

func main() {
	http.HandleFunc("/player", getPlayerHandler)
	port := "8080"
	fmt.Printf("REST API server is listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
	}
}
