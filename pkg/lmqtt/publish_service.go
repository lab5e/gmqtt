package lmqtt

import "github.com/lab5e/lmqtt/pkg/entities"

type publishService struct {
	server *server
}

func (p *publishService) Publish(message *entities.Message) {
	p.server.mu.Lock()
	p.server.deliverMessage("", message, defaultIterateOptions(message.Topic))
	p.server.mu.Unlock()
}
