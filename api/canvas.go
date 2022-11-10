package main

import "log"

type Pixel struct {
    X int `json:"x"`
    Y int `json:"y"`
    Color string `json:"color"`
}

type Canvas struct {
    BufferedPixels []Pixel
    Clients map[*Client]struct{}
    //
    Register chan *Client
    Unregister chan *Client
    Broadcast chan *Pixel
}

func (c *Canvas) HandleEvents() {
    for {
        select {
        case client := <- c.Register:
            log.Println(client)
        case client := <- c.Unregister:
            log.Println(client)
        case pixel := <- c.Broadcast:
            log.Println(pixel)
        }
    }
}

func NewCanvas() *Canvas {
    return &Canvas{
        BufferedPixels: make([]Pixel, 0),
        Clients: make(map[*Client]struct{}),
        Register: make(chan *Client),
        Unregister: make(chan *Client),
        Broadcast: make(chan *Pixel),
    }
}
