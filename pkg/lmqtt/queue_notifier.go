package lmqtt

import (
	"context"

	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/persistence/queue"
)

// queueNotifier implements queue.Notifier interface.
type queueNotifier struct {
	dropHook OnMsgDropped
	sts      *statsManager
	cli      *client
}

// defaultNotifier is used to init the notifier when using a persistent session store which can load session data
// while bootstrapping.
func defaultNotifier(dropHook OnMsgDropped, sts *statsManager, clientID string) *queueNotifier {
	return &queueNotifier{
		dropHook: dropHook,
		sts:      sts,
		cli:      &client{opts: &ClientOptions{ClientID: clientID}, status: Connected + 1},
	}
}

func (q *queueNotifier) notifyDropped(msg *entities.Message, err error) {
	cid := q.cli.opts.ClientID
	q.cli.log("message dropped  client_id=%s: %v", cid, err)
	q.sts.messageDropped(msg.QoS, q.cli.opts.ClientID, err)
	if q.dropHook != nil {
		q.dropHook(context.Background(), cid, msg, err)
	}
}

func (q *queueNotifier) NotifyDropped(elem *queue.Elem, err error) {
	cid := q.cli.opts.ClientID
	if err == queue.ErrDropExpiredInflight && q.cli.IsConnected() {
		q.cli.pl.release(elem.ID())
	}
	if pub, ok := elem.MessageWithID.(*queue.Publish); ok {
		q.notifyDropped(pub.Message, err)
	} else {
		q.cli.log("message dropped  client_ud=%s: %v", cid, err)
	}
}

func (q *queueNotifier) NotifyInflightAdded(delta int) {
	cid := q.cli.opts.ClientID
	if delta > 0 {
		q.sts.addInflight(cid, uint64(delta))
	}
	if delta < 0 {
		q.sts.decInflight(cid, uint64(-delta))
	}

}

func (q *queueNotifier) NotifyMsgQueueAdded(delta int) {
	cid := q.cli.opts.ClientID
	if delta > 0 {
		q.sts.addQueueLen(cid, uint64(delta))
	}
	if delta < 0 {
		q.sts.decQueueLen(cid, uint64(-delta))
	}
}
