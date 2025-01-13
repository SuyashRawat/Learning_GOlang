package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	clientsocket *websocket.Conn
	room         *Room
	sendchan     chan []byte
}

//in the read method
//take an instance that there are 2 users alice and bob
//alice types a message in her browser "hello bob" and sent it via http and it was sent using specialheader which switched the protocal to websocket so basically she sent via websocker
//now i need to forward this message to forwardchan of room in order for the room to broadcast it to all clients
//so we first need to readthe message from the socket and then send to forward channel

// so in this read method message is received by the websocket server and is passed to the forwardchan

// from forwardchan message is send to each clients sendchan
func (c *Client) Read() {

	for {
		_, msg, err := c.clientsocket.ReadMessage()
		if err != nil {
			fmt.Println("failed to retieve message")
			break
		}

		c.room.forwardchan <- msg

	}
	c.clientsocket.Close()

}

// func (c *Client) Read() {
// 	defer func() {
// 		c.room.leave <- c
// 		c.clientsocket.Close()
// 	}()

// 	for {
// 		_, msg, err := c.clientsocket.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Failed to retrieve message:", err)
// 			break
// 		}
// 		c.room.forwardchan <- msg
// 	}
// }

// func(c *Client)Write(){

// 	for msg:=range c.sendchan{
// 		if err:=c.clientsocket.WriteMessage(websocket.TextMessage,msg);err!=nil{
// 			break
// 		}
// 		c.clientsocket.Close()
// 	}
// }

// write method reads from each clients sendchan and the msg is being displayed then
func (c *Client) Write() {
	defer c.clientsocket.Close()

	for msg := range c.sendchan {
		if err := c.clientsocket.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
