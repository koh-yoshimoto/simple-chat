package domain

import (
    "log"

    "github.com/gorilla/websocket"
)

type Client struct {
    ws  *websocket.Conn
    sendCh chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
    return &Client {
        ws: ws,
        sendCh: make(chan []byte),
    }
}

func (c *Client) ReadLoop(broadcast chan<- []byte, unregister chan<- *Client) {
    defer func() {
        c.disconnect(unregister)
    }()

    for  {
        _, jsonMsg, err := c.ws.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("unexpected close error: %v", err)
            }
            break;
        }
        log.Printf("Read: %v", jsonMsg)

        broadcast <- jsonMsg
    }
}

func (c *Client) WriteLoop() {
    defer func() {
        log.Println("WebSocket: Close")
        c.ws.Close()
    }()

    for {
        message := <-c.sendCh

        log.Printf("Write1: %v", message)

        w, err := c.ws.NextWriter(websocket.TextMessage)
        if err != nil {
            log.Printf("NextWriterError: %v", err)
            return
        }

        log.Printf("Write2: %v", message)

        w.Write(message)

        if err := w.Close(); err != nil {
            log.Printf("WriteCloseError: %v", err)
            return
        }
    }
}

func (c *Client) disconnect(unregister chan<- *Client) {
    unregister <- c
    c.ws.Close()
}
