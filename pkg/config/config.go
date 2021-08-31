package config

import (
	"fmt"
)

var (
	defaultPluginConfig = make(map[string]Configuration)
)

// Configuration is the interface that enable the implementation to parse config from the global config file.
// Plugin admin and prometheus are two examples.
type Configuration interface {
	// Validate validates the configuration.
	// If returns error, the broker will not start.
	Validate() error
}

// RegisterDefaultPluginConfig registers the default configuration for the given plugin.
func RegisterDefaultPluginConfig(name string, config Configuration) {
	if _, ok := defaultPluginConfig[name]; ok {
		panic(fmt.Sprintf("duplicated default config for %s plugin", name))
	}
	defaultPluginConfig[name] = config

}

// DefaultConfig return the default configuration.
// If config file is not provided, gmqttd will start with DefaultConfig.
func DefaultConfig() Config {
	c := Config{
		Listeners:         DefaultListeners,
		MQTT:              DefaultMQTTConfig,
		Persistence:       DefaultPersistenceConfig,
		TopicAliasManager: DefaultTopicAliasManager,
	}

	return c
}

// DefaultListeners are the default listeners for the server
var DefaultListeners = []*ListenerConfig{
	{
		Address:    "0.0.0.0:1883",
		TLSOptions: nil,
	},
	{
		Address: "0.0.0.0:8883",
	},
}

// Config is the configration for gmqttd.
type Config struct {
	Listeners         []*ListenerConfig `yaml:"listeners"`
	MQTT              MQTT              `yaml:"mqtt,omitempty"`
	ConfigDir         string            `yaml:"config_dir"`
	Persistence       Persistence       `yaml:"persistence"`
	TopicAliasManager TopicAliasManager `yaml:"topic_alias_manager"`
	DumpPacket        bool              `yaml:"dump_packet"`
}

// TLSOptions are the TLS options for the server
type TLSOptions struct {
	// CACert is the trust CA certificate file.
	CACert string `yaml:"cacert"`
	// Cert is the path to certificate file.
	Cert string `yaml:"cert"`
	// Key is the path to key file.
	Key string `yaml:"key"`
	// Verify indicates whether to verify client cert.
	Verify bool `yaml:"verify"`
}

// ListenerConfig is the (TCP) listener config
type ListenerConfig struct {
	Address     string `yaml:"address"`
	*TLSOptions `yaml:"tls"`
}

// Validate validates the config
func (c Config) Validate() (err error) {
	err = c.MQTT.Validate()
	if err != nil {
		return err
	}
	err = c.Persistence.Validate()
	if err != nil {
		return err
	}
	return nil
}
