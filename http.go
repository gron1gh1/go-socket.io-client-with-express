package main

import (
       "github.com/graarh/golang-socketio"
       "github.com/graarh/golang-socketio/transport"
       "log"
)

type MyEventData struct {
	Data string
}
	
func main() {
       transport := transport.GetDefaultWebsocketTransport()
       ws_url := "ws://localhost/socket.io/?EIO=3&transport=websocket"
       client, err := gosocketio.Dial(ws_url, transport)
       if err != nil {
           log.Fatal(err)
       }
       client.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {
		   log.Println("Connected!")
		   client.Emit("msg", MyEventData{"Hello"})
       })

	   // Block to give client time to connect 
	   select
		{

		}
       client.Close()
}