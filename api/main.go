package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type App struct {
    Canvas *Canvas
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func (a *App) serveWs(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade failed")
        return
    }

    client := &Client{
        Conn: c,
        Canvas: a.Canvas,
    }

    go client.ReadPump()
}

func main() {
    canvas := NewCanvas()

    app := &App{
        Canvas: canvas,
    }

    go canvas.HandleEvents()

    http.HandleFunc("/ws", app.serveWs)
    http.ListenAndServe(":8080", nil)
}
