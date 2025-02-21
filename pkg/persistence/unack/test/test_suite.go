package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/packets"
	"github.com/lab5e/lmqtt/pkg/persistence/unack"
)

var (
	// ServerConfig is the default server config used in the tests
	ServerConfig = config.Config{}
	// ClientID is the client ID used in the tests
	ClientID = "cid"
)

// Suite runs tests on an unack.Store implementation
func Suite(t *testing.T, store unack.Store) {
	a := assert.New(t)
	a.Nil(store.Init(false))
	for i := packets.PacketID(1); i < 10; i++ {
		rs, err := store.Set(i)
		a.Nil(err)
		a.False(rs)
		rs, err = store.Set(i)
		a.Nil(err)
		a.True(rs)
		err = store.Remove(i)
		a.Nil(err)
		rs, err = store.Set(i)
		a.Nil(err)
		a.False(rs)

	}
	a.Nil(store.Init(false))
	for i := packets.PacketID(1); i < 10; i++ {
		rs, err := store.Set(i)
		a.Nil(err)
		a.True(rs)
		err = store.Remove(i)
		a.Nil(err)
		rs, err = store.Set(i)
		a.Nil(err)
		a.False(rs)
	}
	a.Nil(store.Init(true))
	for i := packets.PacketID(1); i < 10; i++ {
		rs, err := store.Set(i)
		a.Nil(err)
		a.False(rs)
	}

}
