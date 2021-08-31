package mem

import (
	"sync"

	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/persistence/session"
)

var _ session.Store = (*Store)(nil)

// New creates a memory-backed session store
func New() *Store {
	return &Store{
		mu:   sync.Mutex{},
		sess: make(map[string]*entities.Session),
	}
}

// Store is the memory store type
type Store struct {
	mu   sync.Mutex
	sess map[string]*entities.Session
}

// Set sets a new session in the store
func (s *Store) Set(session *entities.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sess[session.ClientID] = session
	return nil
}

// Remove removes a session from the store
func (s *Store) Remove(clientID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sess, clientID)
	return nil
}

// Get returns a session from the store for the specified client
func (s *Store) Get(clientID string) (*entities.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sess[clientID], nil
}

// GetAll returns all sessions. This returns nil for the memory store
func (s *Store) GetAll() ([]*entities.Session, error) {
	return nil, nil
}

// SetSessionExpiry sets the expiration time for the sessions
func (s *Store) SetSessionExpiry(clientID string, expiry uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s, ok := s.sess[clientID]; ok {
		s.ExpiryInterval = expiry

	}
	return nil
}

// Iterate iterates over the sessions in the store
func (s *Store) Iterate(fn session.IterateFn) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.sess {
		cont := fn(v)
		if !cont {
			break
		}
	}
	return nil
}
