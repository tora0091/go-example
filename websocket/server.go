package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcase = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
	Message string `json:message`
}

func init() {
	log.SetPrefix("[server] ")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ws", handleClient)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("access from %s", r.RemoteAddr)
	http.ServeFile(w, r, "views/index.html")
}

func handleClient(w http.ResponseWriter, r *http.Request) {
	go broadcastMessagesToClients()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	clients[conn] = true

	broadcastMessagesFromClient(conn)
}

func broadcastMessagesToClients() {
	for {
		message := <-broadcase
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func broadcastMessagesFromClient(conn *websocket.Conn) {
	for {
		var message Message

		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			break
		}
		broadcase <- message
	}
}
