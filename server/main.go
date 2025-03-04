package main

import (
	"fmt"
	"net/http"

	"smashrift.mikekaipis.com/api"
	"smashrift.mikekaipis.com/benchmarking"
	"smashrift.mikekaipis.com/state"
)

func main() {
	benchmarking.Init()
	state.Init()

	http.HandleFunc("/ws", api.WsHandler)

	go api.HandleMessages()
	go api.SendTickerMessages()

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
