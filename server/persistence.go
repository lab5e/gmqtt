package server

import (
	"github.com/lab5e/gmqtt/config"
	"github.com/lab5e/gmqtt/persistence/queue"
	"github.com/lab5e/gmqtt/persistence/session"
	"github.com/lab5e/gmqtt/persistence/subscription"
	"github.com/lab5e/gmqtt/persistence/unack"
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
