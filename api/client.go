package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	Canvas *Canvas
}

// TODO: PING PONG DEAD CLIENT CHECK

func (c *Client) ReadPump() {
	var pixel *Pixel
	for {
		err := c.Conn.ReadJSON(&pixel)
		if err != nil {
			log.Println(err)
			return
		}

		c.Canvas.Broadcast <- pixel
	}
}

func (c *Client) WritePump() {

}
