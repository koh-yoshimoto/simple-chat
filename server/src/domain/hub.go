package domain

import (
    "context"

    "github.com/koh-yoshimoto/simple-chat/src/services"
)

type Hub struct {
    Clients map[*Client]bool
    RegisterCh chan *Client
    UnRegisterCh chan *Client
    BroadcastCh chan []byte
    pubsub *services.PubSubService
}

const broadcastChan = "braodcast"

func NewHub(pubsub *services.PubSubService) *Hub {
    return &Hub {
        Clients: make(map[*Client]bool),
        RegisterCh: make(chan *Client),
        UnRegisterCh: make(chan *Client),
        BroadcastCh: make(chan []byte),
        pubsub: pubsub,
    }
}

func (h *Hub) RunLoop() {
    for {
        select {
        case client := <-h.RegisterCh:
            h.register(client)

        case client := <-h.UnRegisterCh:
            h.unregister(client)

        case msg := <-h.BroadcastCh:
            // h.broadcastToAllClient(msg)
            h.publishMessage(msg);
        }
    }
}

func (h *Hub) SubscribeMessages() {
    ch := h.pubsub.Subscribe(context.TODO(), broadcastChan)
    for msg := range ch {
        h.broadcastToAllClient([]byte(msg.Payload))
    }
}

func (h *Hub) publishMessage(msg []byte) {
    h.pubsub.Publish(context.TODO(), broadcastChan, msg)
}

func (h *Hub) register(c *Client) {
    h.Clients[c] = true
}

func (h *Hub) unregister(c *Client) {
    delete(h.Clients, c)
}

func (h *Hub) broadcastToAllClient(msg []byte) {
    for c := range h.Clients {
        c.sendCh <- msg
    }
}

