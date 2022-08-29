package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/koh-yoshimoto/simple-chat/src/domain"
	"github.com/koh-yoshimoto/simple-chat/src/handlers"
	"github.com/koh-yoshimoto/simple-chat/src/services"
)

func main() {
    pubsub := services.NewPubSubService()
    hub := domain.NewHub(pubsub)
    go hub.SubscribeMessages()
    go hub.RunLoop()

    http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

    port := "80"
    log.Printf("Listening on port %s", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
        log.Panicln("Serve Error:", err);
    }
}

