package models

import (
	"chat_application/trace"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte
	// join is a channel for clients wishing to join the room.
	join chan *Client
	// leave is a channel for clients wishing to leave the room.
	leave chan *Client
	// clients holds all current clients in this room.
	clients map[*Client]bool
	// tracer will receive trace information of activity
	// in the room.
	Tracer trace.Tracer
}

func NewRoom() *Room {
	return &Room{
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		forward: make(chan []byte),
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			fmt.Println("welcome client")
			r.clients[client] = true
			r.Tracer.Trace("New client joined")

		case client := <-r.leave:
			fmt.Println("client left")
			delete(r.clients, client)
			close(client.sendchan)
		case msg := <-r.forward:
			fmt.Println("forwarding the message")
			for client := range r.clients {
				select {
				case client.sendchan <- msg:
					fmt.Println("sent to client")
					r.Tracer.Trace(" -- sent to client")
				default:
					delete(r.clients, client)
					close(client.sendchan)
					r.Tracer.Trace(" -- failed to send, cleaned up client")
				}
			}
		}
	}
}

func (r *Room) ServeNewRoom(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	client := &Client{
		clientsocket: socket,
		sendchan:     make(chan []byte, messageBufferSize),
		room:         r,
	}
	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.Write()
	client.Read()
}
