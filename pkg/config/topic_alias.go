package config

// TopicAliasType is a type for topic alias factories
type TopicAliasType = string

// Topic alias type managerss
const (
	TopicAliasMgrTypeFIFO TopicAliasType = "fifo"
)

var (
	// DefaultTopicAliasManager is the default value of TopicAliasManager
	DefaultTopicAliasManager = TopicAliasManager{
		Type: TopicAliasMgrTypeFIFO,
	}
)

// TopicAliasManager is the config of the topic alias manager.
type TopicAliasManager struct {
	Type TopicAliasType
}
