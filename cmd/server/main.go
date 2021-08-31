package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/lmqtt"

	_ "github.com/lab5e/lmqtt/pkg/persistence"     // Default store is memory
	_ "github.com/lab5e/lmqtt/pkg/topicalias/fifo" // Message handling
)

type params struct {
	ListenAddress    string
	GenerateMessages bool
	Topic            string
	Username         string
	Password         string
}

var cfg params

func main() {

	flag.StringVar(&cfg.ListenAddress, "listen-address", ":1883", "Listen address for MQTT server")
	flag.BoolVar(&cfg.GenerateMessages, "generate", false, "Generate messages")
	flag.StringVar(&cfg.Topic, "topic", "#/generated", "Topic for generated messages")
	flag.StringVar(&cfg.Username, "username", "", "Optional username")
	flag.StringVar(&cfg.Password, "password", "", "Optional password")

	flag.Parse()

	listener, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		log.Fatalf("Could not create listener: %v", err)
	}

	config := config.DefaultConfig()

	server := lmqtt.New(
		lmqtt.WithTCPListener(listener),
		lmqtt.WithConfig(config),
		lmqtt.WithHook(lmqtt.Hooks{
			OnAccept:            onAccept,
			OnStop:              onStop,
			OnSubscribe:         onSubscribe,
			OnSubscribed:        onSubscribed,
			OnUnsubscribe:       onUnsubscribe,
			OnUnsubscribed:      onUnsubscribed,
			OnMsgArrived:        onMsgArrived,
			OnBasicAuth:         onBasicAuth,
			OnEnhancedAuth:      onEnhancedAuth,
			OnReAuth:            onReauth,
			OnConnected:         onConnected,
			OnSessionCreated:    onSessionCreated,
			OnSessionResumed:    onSessionResumed,
			OnSessionTerminated: onSessionTerminated,
			OnDelivered:         onDelivered,
			OnClosed:            onClosed,
			OnMsgDropped:        onMsgDropped,
			OnWillPublish:       onWillPublish,
			OnWillPublished:     onWillPublished,
		}))

	if cfg.GenerateMessages {
		go func(p lmqtt.Publisher) {
			if cfg.Topic == "" {
				log.Fatalf("Needs a topic when generating messages")
			}
			msgNo := 1
			for {
				p.Publish(&entities.Message{
					Topic:   cfg.Topic,
					Payload: []byte(fmt.Sprintf("This is message %d", msgNo)),
					QoS:     byte(1),
				})

				msgNo++
				time.Sleep(1 * time.Second)
			}
		}(server.Publisher())
	}

	if err := server.Run(); err != nil {
		log.Printf("Error running server: %v", err)
	}
}
