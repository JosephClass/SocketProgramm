package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

type M map[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("Index.html")
		if err != nil {
			http.Error(w,"Error Message: Could not open the requested file", http.StatusInternalServerError)
			return 
		}

		fmt.Fprintf(w,"%s", content)
	})
}
