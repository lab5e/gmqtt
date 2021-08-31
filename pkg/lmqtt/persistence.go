package lmqtt

import (
	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/persistence/queue"
	"github.com/lab5e/lmqtt/pkg/persistence/session"
	"github.com/lab5e/lmqtt/pkg/persistence/subscription"
	"github.com/lab5e/lmqtt/pkg/persistence/unack"
)

// NewPersistence creates a new persistence layer
type NewPersistence func(config config.Config) (Persistence, error)

// Persistence is the storage layer
type Persistence interface {
	Open() error
	NewQueueStore(config config.Config, defaultNotifier queue.Notifier, clientID string) (queue.Store, error)
	NewSubscriptionStore(config config.Config) (subscription.Store, error)
	NewSessionStore(config config.Config) (session.Store, error)
	NewUnackStore(config config.Config, clientID string) (unack.Store, error)
	Close() error
}
