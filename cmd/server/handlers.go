package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/lmqtt"
	"github.com/lab5e/lmqtt/pkg/packets"
)

// This is stub handlers for the simple server.

func onAccept(ctx context.Context, conn net.Conn) bool {
	log.Printf("onAccept: %s", conn.RemoteAddr())
	return true
}

func onStop(ctx context.Context) {
	log.Printf("onStop")
}

func onSubscribe(ctx context.Context, client lmqtt.Client, req *lmqtt.SubscribeRequest) error {
	log.Printf("onSubscribe[%s]", client.ClientOptions().ClientID)
	if req.Subscribe != nil {
		for _, v := range req.Subscribe.Topics {
			// This is a test case for the paho.mqtt tests. If the client attempts
			// to subscribe to this topic it should fail.
			if v.Name == "test/nosubscribe" {
				return errors.New("invalid topic")
			}
		}
	}
	return nil
}

func onSubscribed(ctx context.Context, client lmqtt.Client, sub *entities.Subscription) {
	log.Printf("onSubscribed[%s]: topic_filter=%+v", client.ClientOptions().ClientID, sub.TopicFilter)
}

func onUnsubscribe(ctx context.Context, client lmqtt.Client, req *lmqtt.UnsubscribeRequest) error {
	log.Printf("onUnsubscribe[%s]", client.ClientOptions().ClientID)
	return nil
}

func onUnsubscribed(ctx context.Context, client lmqtt.Client, topicName string) {
	log.Printf("onUnsubscribed[%s] topic=%+v", client.ClientOptions().ClientID, topicName)
}

func onMsgArrived(ctx context.Context, client lmqtt.Client, req *lmqtt.MsgArrivedRequest) error {
	log.Printf("onMsgArrived[%s] format=%+v, payload=%d bytes", client.ClientOptions().ClientID, req.Message.PayloadFormat, len(req.Message.Payload))
	return nil
}

func onBasicAuth(ctx context.Context, client lmqtt.Client, req *lmqtt.ConnectRequest) error {
	log.Printf("onBasicAuth[%s] username=%s, password=%s", client.ClientOptions().ClientID, string(req.Connect.Username), string(req.Connect.Password))
	if cfg.Username != "" && cfg.Password != "" {
		if cfg.Username != string(req.Connect.Username) ||
			cfg.Password != string(req.Connect.Password) {
			return errors.New("unauthenticated")
		}
	}
	return nil
}

func onEnhancedAuth(ctx context.Context, client lmqtt.Client, req *lmqtt.ConnectRequest) (*lmqtt.EnhancedAuthResponse, error) {
	log.Printf("onEnhancedAuth[%s] req=%+v", client.ClientOptions().ClientID, req)
	return nil, nil
}

func onReauth(ctx context.Context, client lmqtt.Client, auth *packets.Auth) (*lmqtt.AuthResponse, error) {
	log.Printf("onReauth[%s] req=%+v", client.ClientOptions().ClientID, auth)
	return nil, nil
}

func onConnected(ctx context.Context, client lmqtt.Client) {
	log.Printf("onConnected[%s]", client.ClientOptions().ClientID)
}

func onSessionCreated(ctx context.Context, client lmqtt.Client) {
	log.Printf("onSessionCreated[%s]", client.ClientOptions().ClientID)
}

func onSessionResumed(ctx context.Context, client lmqtt.Client) {
	log.Printf("onSessionResumed[%s]", client.ClientOptions().ClientID)
}

func onSessionTerminated(ctx context.Context, clientID string, reason lmqtt.SessionTerminatedReason) {
	log.Printf("onSessionTerminated[%s] reason=%+v", clientID, reason)
}

func onDelivered(ctx context.Context, client lmqtt.Client, msg *entities.Message) {
	log.Printf("onDelivered[%s] bytes=%d", client.ClientOptions().ClientID, len(msg.Payload))
}

func onMsgDropped(ctx context.Context, clientID string, msg *entities.Message, err error) {
	log.Printf("onMsgDropped[%s] bytes=%d err=%v", clientID, len(msg.Payload), err)
}

func onWillPublish(ctx context.Context, clientID string, req *lmqtt.WillMsgRequest) {
	log.Printf("onWillPublish[%s] bytes=%d", clientID, len(req.Message.Payload))
}

func onWillPublished(ctx context.Context, clientID string, req *entities.Message) {
	log.Printf("onWillPublish[%s] bytes=%d", clientID, len(req.Payload))
}

func onClosed(ctx context.Context, client lmqtt.Client, err error) {
	log.Printf("onClosed[%s] err=%v", client.ClientOptions().ClientID, err)
}
