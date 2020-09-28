package main

import (
	"log"
	"net/http"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	//handle connected
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		//join them to room
		c.Join("chat")
	})

	type Message struct {
		Name string `json:"name"`
		Message string `json:"message"`
	}

	//handle custom event
	server.On("msg", func(c *gosocketio.Channel, msg Message) string {
		//send event to all in room
		log.Println(msg)
		return "OK"
	})
	channel, _ := server.GetChannel("msg")

	channel.Emit("msg","my data")
	//setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	serveMux.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Panic(http.ListenAndServe(":80", serveMux))
}