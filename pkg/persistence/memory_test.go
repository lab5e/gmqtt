package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/lmqtt"
	queue_test "github.com/lab5e/lmqtt/pkg/persistence/queue/test"
	sess_test "github.com/lab5e/lmqtt/pkg/persistence/session/test"
	"github.com/lab5e/lmqtt/pkg/persistence/subscription"
	sub_test "github.com/lab5e/lmqtt/pkg/persistence/subscription/test"
	unack_test "github.com/lab5e/lmqtt/pkg/persistence/unack/test"
)

type MemorySuite struct {
	suite.Suite
	p lmqtt.Persistence
}

func (s *MemorySuite) TestQueue() {
	a := assert.New(s.T())
	qs, err := s.p.NewQueueStore(queue_test.ServerConfig, queue_test.Notifier, queue_test.ClientID)
	a.Nil(err)
	queue_test.TestQueue(s.T(), qs)
}
func (s *MemorySuite) TestSubscription() {
	newFn := func() subscription.Store {
		st, err := s.p.NewSubscriptionStore(queue_test.ServerConfig)
		if err != nil {
			panic(err)
		}
		return st
	}
	sub_test.Suite(s.T(), newFn)
}

func (s *MemorySuite) TestSession() {
	a := assert.New(s.T())
	st, err := s.p.NewSessionStore(queue_test.ServerConfig)
	a.Nil(err)
	sess_test.Suite(s.T(), st)
}

func (s *MemorySuite) TestUnack() {
	a := assert.New(s.T())
	st, err := s.p.NewUnackStore(unack_test.ServerConfig, unack_test.ClientID)
	a.Nil(err)
	unack_test.Suite(s.T(), st)
}

func TestMemory(t *testing.T) {
	p, err := NewMemory(config.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	suite.Run(t, &MemorySuite{
		p: p,
	})
}
