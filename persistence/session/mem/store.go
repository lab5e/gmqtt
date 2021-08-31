package mem

import (
	"sync"

	"github.com/lab5e/lmqtt/persistence/session"
	"github.com/lab5e/lmqtt/pkg/entities"
)

var _ session.Store = (*Store)(nil)

func New() *Store {
	return &Store{
		mu:   sync.Mutex{},
		sess: make(map[string]*entities.Session),
	}
}

type Store struct {
	mu   sync.Mutex
	sess map[string]*entities.Session
}

func (s *Store) Set(session *entities.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sess[session.ClientID] = session
	return nil
}

func (s *Store) Remove(clientID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sess, clientID)
	return nil
}

func (s *Store) Get(clientID string) (*entities.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sess[clientID], nil
}

func (s *Store) GetAll() ([]*entities.Session, error) {
	return nil, nil
}

func (s *Store) SetSessionExpiry(clientID string, expiry uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s, ok := s.sess[clientID]; ok {
		s.ExpiryInterval = expiry

	}
	return nil
}

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
