package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"smashrift.mikekaipis.com/state"
	"smashrift.mikekaipis.com/utils"
)

type Message struct {
	Type    string      `json:"type"` // "chat" or "update" or "player"
	Content interface{} `json:"content"`
}

var clients = make(map[*websocket.Conn]string)
var chatBroadcast = make(chan Message)
var updateBroadcast = make(chan Message)
var mutex = &sync.Mutex{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "<http://yourdomain.com>" || true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer func() {
		mutex.Lock()
		playerID, exists := clients[conn]
		if exists {
			delete(clients, conn) // Remove client from map
			mutex.Unlock()
			state.DeletePlayer(playerID) // Delete player from state after unlocking
		} else {
			mutex.Unlock()
		}

		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}()

	mutex.Lock()
	clients[conn] = utils.RandSeq(6) // Random identifier for client
	mutex.Unlock()

	// add client to state
	player := state.Player{
		ID:        clients[conn],
		Name:      clients[conn],
		Position:  state.GetRandMapPosition(),
		Direction: utils.GetRandDirection(),
		Speed:     2,
	}
	state.AddPlayer(player)

	// Send player ID to client
	playerMessage := Message{
		Type:    "player",
		Content: player,
	}
	playerMessageJSON, err := json.Marshal(playerMessage)
	if err != nil {
		fmt.Println("Error marshaling player message:", err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, playerMessageJSON)
	if err != nil {
		fmt.Println("Error writing message:", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println("Error unmarshaling message:", err)
			continue
		}

		if msg.Type == "chat" {
			chatMsg := Message{
				Type:    "chat",
				Content: msg.Content,
			}
			chatBroadcast <- chatMsg
		} else if msg.Type == "update-direction" {
			fmt.Println("Update direction message received", msg)
			updateDirectionMessage := utils.Vector2{
				X: msg.Content.(map[string]interface{})["x"].(float64),
				Y: msg.Content.(map[string]interface{})["y"].(float64),
			}
			ID := clients[conn]
			state.UpdatePlayer(ID, updateDirectionMessage.Normalize(), nil)
		}
	}
}

// HandleMessages sends chat messages to all clients
func HandleMessages() {
	for {
		msg := <-chatBroadcast
		if msg.Type == "chat" {
			sendToAllClients(msg)
		}
		// else if msg.Type == "update-direction" {
		// 	updateDirectionMessage := msg.Content.(UpdateDirectionMessage)
		// 	state.UpdatePlayer(updateDirectionMessage.ID, updateDirectionMessage.Content.Normalize(), nil)
		// }
	}
}

func SendTickerMessages() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	updateCount := 0

	for {
		<-ticker.C // Now this properly waits for 10ms before each loop

		updateCount++

		players := state.GetPlayers()

		msg := Message{
			Type:    "update",
			Content: players,
		}
		sendToAllClients(msg)
	}
}

// Sends messages to all connected clients
func sendToAllClients(msg Message) {
	jsonMessage, _ := json.Marshal(msg)

	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
