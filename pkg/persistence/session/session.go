package session

import "github.com/lab5e/lmqtt/pkg/entities"

// IterateFn is the callback function used by Iterate()
// Return false means to stop the iteration.
type IterateFn func(session *entities.Session) bool

// Store is the session store interface
type Store interface {
	Set(session *entities.Session) error
	Remove(clientID string) error
	Get(clientID string) (*entities.Session, error)
	Iterate(fn IterateFn) error
	SetSessionExpiry(clientID string, expiry uint32) error
}
