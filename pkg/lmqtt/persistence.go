package lmqtt

import (
	"github.com/lab5e/lmqtt/persistence/queue"
	"github.com/lab5e/lmqtt/persistence/session"
	"github.com/lab5e/lmqtt/persistence/subscription"
	"github.com/lab5e/lmqtt/persistence/unack"
	"github.com/lab5e/lmqtt/pkg/config"
)

type NewPersistence func(config config.Config) (Persistence, error)

type Persistence interface {
	Open() error
	NewQueueStore(config config.Config, defaultNotifier queue.Notifier, clientID string) (queue.Store, error)
	NewSubscriptionStore(config config.Config) (subscription.Store, error)
	NewSessionStore(config config.Config) (session.Store, error)
	NewUnackStore(config config.Config, clientID string) (unack.Store, error)
	Close() error
}
