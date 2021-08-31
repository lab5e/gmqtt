package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/lab5e/lmqtt/config"
	queue_test "github.com/lab5e/lmqtt/persistence/queue/test"
	sess_test "github.com/lab5e/lmqtt/persistence/session/test"
	"github.com/lab5e/lmqtt/persistence/subscription"
	sub_test "github.com/lab5e/lmqtt/persistence/subscription/test"
	unack_test "github.com/lab5e/lmqtt/persistence/unack/test"
	"github.com/lab5e/lmqtt/pkg/lmqtt"
)

type MemorySuite struct {
	suite.Suite
	p lmqtt.Persistence
}

func (s *MemorySuite) TestQueue() {
	a := assert.New(s.T())
	qs, err := s.p.NewQueueStore(queue_test.TestServerConfig, queue_test.TestNotifier, queue_test.TestClientID)
	a.Nil(err)
	queue_test.TestQueue(s.T(), qs)
}
func (s *MemorySuite) TestSubscription() {
	newFn := func() subscription.Store {
		st, err := s.p.NewSubscriptionStore(queue_test.TestServerConfig)
		if err != nil {
			panic(err)
		}
		return st
	}
	sub_test.TestSuite(s.T(), newFn)
}

func (s *MemorySuite) TestSession() {
	a := assert.New(s.T())
	st, err := s.p.NewSessionStore(queue_test.TestServerConfig)
	a.Nil(err)
	sess_test.TestSuite(s.T(), st)
}

func (s *MemorySuite) TestUnack() {
	a := assert.New(s.T())
	st, err := s.p.NewUnackStore(unack_test.TestServerConfig, unack_test.TestClientID)
	a.Nil(err)
	unack_test.TestSuite(s.T(), st)
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
