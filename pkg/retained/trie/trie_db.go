package trie

import (
	"sync"

	"github.com/lab5e/lmqtt/pkg/entities"
	"github.com/lab5e/lmqtt/pkg/retained"
)

// trieDB implement the retain.Store, it use trie tree  to store retain messages .
type trieDB struct {
	sync.RWMutex
	userTrie   *topicTrie
	systemTrie *topicTrie
}

func (t *trieDB) Iterate(fn retained.IterateFn) {
	t.RLock()
	defer t.RUnlock()
	if !t.userTrie.preOrderTraverse(fn) {
		return
	}
	t.systemTrie.preOrderTraverse(fn)
}

func (t *trieDB) getTrie(topicName string) *topicTrie {
	if isSystemTopic(topicName) {
		return t.systemTrie
	}
	return t.userTrie
}

// GetRetainedMessage return the retain message of the given topic name.
// return nil if the topic name not exists
func (t *trieDB) GetRetainedMessage(topicName string) *entities.Message {
	t.RLock()
	defer t.RUnlock()
	node := t.getTrie(topicName).find(topicName)
	if node != nil {
		return node.msg.Copy()
	}
	return nil
}

// ClearAll clear all retain messages.
func (t *trieDB) ClearAll() {
	t.Lock()
	defer t.Unlock()
	t.systemTrie = newTopicTrie()
	t.userTrie = newTopicTrie()
}

// AddOrReplace add or replace a retain message.
func (t *trieDB) AddOrReplace(message *entities.Message) {
	t.Lock()
	defer t.Unlock()
	t.getTrie(message.Topic).addRetainMsg(message.Topic, message)
}

// remove remove the retain message of the topic name.
func (t *trieDB) Remove(topicName string) {
	t.Lock()
	defer t.Unlock()
	t.getTrie(topicName).remove(topicName)
}

// GetMatchedMessages returns all messages that match the topic filter.
func (t *trieDB) GetMatchedMessages(topicFilter string) []*entities.Message {
	t.RLock()
	defer t.RUnlock()
	return t.getTrie(topicFilter).getMatchedMessages(topicFilter)
}

// NewStore creates a new Store using a trie tree
func NewStore() retained.Store {
	return &trieDB{
		userTrie:   newTopicTrie(),
		systemTrie: newTopicTrie(),
	}
}
