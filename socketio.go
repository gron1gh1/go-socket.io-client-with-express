package main

import (
	"log"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type MyEventData struct {
	Data int
}

func main() {
	transport := transport.GetDefaultWebsocketTransport()
	ws_url := "ws://localhost/socket.io/?EIO=3&transport=websocket"
	client, err := gosocketio.Dial(ws_url, transport)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	client.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {
		log.Println("Connected!")

	})
	go func() {
		i := 0
		for {
			i++
			client.Emit("data", MyEventData{i})
			time.Sleep(time.Millisecond * 500)

			log.Println(i)
		}
	}()
	// Block to give client time to connect
	select {}
}
