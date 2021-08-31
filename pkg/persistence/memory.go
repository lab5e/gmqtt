package persistence

import (
	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/lmqtt"
	"github.com/lab5e/lmqtt/pkg/persistence/queue"
	mem_queue "github.com/lab5e/lmqtt/pkg/persistence/queue/mem"
	"github.com/lab5e/lmqtt/pkg/persistence/session"
	mem_session "github.com/lab5e/lmqtt/pkg/persistence/session/mem"
	"github.com/lab5e/lmqtt/pkg/persistence/subscription"
	mem_sub "github.com/lab5e/lmqtt/pkg/persistence/subscription/mem"
	"github.com/lab5e/lmqtt/pkg/persistence/unack"
	mem_unack "github.com/lab5e/lmqtt/pkg/persistence/unack/mem"
)

func init() {
	lmqtt.RegisterPersistenceFactory("memory", NewMemory)
}

func NewMemory(config config.Config) (lmqtt.Persistence, error) {
	return &memory{}, nil
}

type memory struct {
}

func (m *memory) NewUnackStore(config config.Config, clientID string) (unack.Store, error) {
	return mem_unack.New(mem_unack.Options{
		ClientID: clientID,
	}), nil
}

func (m *memory) NewSessionStore(config config.Config) (session.Store, error) {
	return mem_session.New(), nil
}

func (m *memory) Open() error {
	return nil
}
func (m *memory) NewQueueStore(config config.Config, defaultNotifier queue.Notifier, clientID string) (queue.Store, error) {
	return mem_queue.New(mem_queue.Options{
		MaxQueuedMsg:    config.MQTT.MaxQueuedMsg,
		InflightExpiry:  config.MQTT.InflightExpiry,
		ClientID:        clientID,
		DefaultNotifier: defaultNotifier,
	})
}

func (m *memory) NewSubscriptionStore(config config.Config) (subscription.Store, error) {
	return mem_sub.NewStore(), nil
}

func (m *memory) Close() error {
	return nil
}
