package lmqtt

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/lab5e/lmqtt/pkg/config"
	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/packets"
	"github.com/lab5e/lmqtt/pkg/persistence/queue"
	"github.com/lab5e/lmqtt/pkg/persistence/subscription/mem"
)

type testDeliverMsg struct {
	srv *server
}

func newTestDeliverMsg(ctrl *gomock.Controller, subscriber string) *testDeliverMsg {
	sub := mem.NewStore()
	srv := &server{
		subscriptionsDB: sub,
		queueStore:      make(map[string]queue.Store),
		config:          config.DefaultConfig(),
		statsManager:    newStatsManager(sub),
	}
	mockQueue := queue.NewMockStore(ctrl)
	srv.queueStore[subscriber] = mockQueue
	return &testDeliverMsg{
		srv: srv,
	}
}

func TestServer_deliverMessage(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subscriber := "subCli"
	ts := newTestDeliverMsg(ctrl, subscriber)
	srcCli := "srcCli"
	msg := &entities.Message{
		Topic:   "/abc",
		Payload: []byte("abc"),
		QoS:     2,
	}
	srv := ts.srv
	_, _ = srv.subscriptionsDB.Subscribe(subscriber, &entities.Subscription{
		ShareName:   "",
		TopicFilter: "/abc",
		QoS:         1,
	}, &entities.Subscription{
		ShareName:   "",
		TopicFilter: "/+",
		QoS:         2,
	})

	mockQueue := srv.queueStore[subscriber].(*queue.MockStore)
	// test only once
	srv.config.MQTT.DeliveryMode = OnlyOnce
	mockQueue.EXPECT().Add(gomock.Any()).Do(func(elem *queue.Elem) {
		a.EqualValues(elem.MessageWithID.(*queue.Publish).QoS, 2)
	})

	a.True(srv.deliverMessage(srcCli, msg, defaultIterateOptions(msg.Topic)))

	// test overlap
	srv.config.MQTT.DeliveryMode = Overlap
	qos := map[byte]int{
		packets.Qos1: 0,
		packets.Qos2: 0,
	}
	mockQueue.EXPECT().Add(gomock.Any()).Do(func(elem *queue.Elem) {
		_, ok := qos[elem.MessageWithID.(*queue.Publish).QoS]
		a.True(ok)
		qos[elem.MessageWithID.(*queue.Publish).QoS]++
	}).Times(2)

	a.True(srv.deliverMessage(srcCli, msg, defaultIterateOptions(msg.Topic)))

	a.Equal(1, qos[packets.Qos1])
	a.Equal(1, qos[packets.Qos2])

	msg = &entities.Message{
		Topic: "abcd",
	}
	a.False(srv.deliverMessage(srcCli, msg, defaultIterateOptions(msg.Topic)))

}

func TestServer_deliverMessage_sharedSubscription(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subscriber := "subCli"
	ts := newTestDeliverMsg(ctrl, subscriber)
	srcCli := "srcCli"
	msg := &entities.Message{
		Topic:   "/abc",
		Payload: []byte("abc"),
		QoS:     2,
	}
	srv := ts.srv
	// add 2 shared and 2 non-shared subscription which both match the message topic: /abc
	_, _ = srv.subscriptionsDB.Subscribe(subscriber, &entities.Subscription{
		ShareName:   "abc",
		TopicFilter: "/abc",
		QoS:         1,
	}, &entities.Subscription{
		ShareName:   "abc",
		TopicFilter: "/+",
		QoS:         2,
	}, &entities.Subscription{
		TopicFilter: "#",
		QoS:         2,
	}, &entities.Subscription{
		TopicFilter: "/abc",
		QoS:         1,
	})

	mockQueue := srv.queueStore[subscriber].(*queue.MockStore)
	// test only once
	qos := map[byte]int{
		packets.Qos1: 0,
		packets.Qos2: 0,
	}
	srv.config.MQTT.DeliveryMode = OnlyOnce
	mockQueue.EXPECT().Add(gomock.Any()).Do(func(elem *queue.Elem) {
		_, ok := qos[elem.MessageWithID.(*queue.Publish).QoS]
		a.True(ok)
		qos[elem.MessageWithID.(*queue.Publish).QoS]++

	}).Times(3)

	a.True(srv.deliverMessage(srcCli, msg, defaultIterateOptions(msg.Topic)))
	a.Equal(1, qos[packets.Qos1])
	a.Equal(2, qos[packets.Qos2])

	// test overlap
	srv.config.MQTT.DeliveryMode = Overlap
	qos = map[byte]int{
		packets.Qos1: 0,
		packets.Qos2: 0,
	}
	mockQueue.EXPECT().Add(gomock.Any()).Do(func(elem *queue.Elem) {
		_, ok := qos[elem.MessageWithID.(*queue.Publish).QoS]
		a.True(ok)
		qos[elem.MessageWithID.(*queue.Publish).QoS]++
	}).Times(4)
	a.True(srv.deliverMessage(srcCli, msg, defaultIterateOptions(msg.Topic)))
	a.Equal(2, qos[packets.Qos1])
	a.Equal(2, qos[packets.Qos2])

}
