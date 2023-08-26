package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func handleWebSocket(serverObj *server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}

	serverObj.mu.Lock()
	userName := serverObj.userHashes[string(message)]
	if userName == "" {
		fmt.Println("Error finding user:", err)
		serverObj.mu.Unlock()
		return
	}
	serverObj.mu.Unlock()



	log.Println("WebSocket Client connected " + userName)

	for {
		serverObj.mu.Lock()
		if len(serverObj.waitingRoom) > 1 {
			// Send a message to the client
			err = conn.WriteMessage(websocket.TextMessage, []byte("Opponent Found"))
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
			serverObj.mu.Unlock()
			newGameBoard := MancalaGameBoard{
				p1Name: serverObj.waitingRoom[0],
				p2Name: serverObj.waitingRoom[1],
			}
			newGameBoard.resetBoard()
			serverObj.gamesInProgress = append(serverObj.gamesInProgress, newGameBoard)
			if len(serverObj.waitingRoom) > 2 {
				serverObj.waitingRoom = serverObj.waitingRoom[2:]
			} else {
				serverObj.waitingRoom = []string{}
			}

			return
		}
		serverObj.mu.Unlock()
	}
}


func websocketSetup(serverObj *server) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(serverObj, w, r)
	})
	log.Printf("Websocket listening at %v", *wsport)
	http.ListenAndServe(":"+strconv.Itoa(*wsport), nil)
}