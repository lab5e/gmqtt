package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
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
	// Unmarshaler defined how to unmarshal YAML into the config structure.
	yaml.Unmarshaler
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
		Listeners: DefaultListeners,
		MQTT:      DefaultMQTTConfig,
		Log: LogConfig{
			Level:  "info",
			Format: "text",
		},
		Persistence:       DefaultPersistenceConfig,
		TopicAliasManager: DefaultTopicAliasManager,
	}

	return c
}

var DefaultListeners = []*ListenerConfig{
	{
		Address:    "0.0.0.0:1883",
		TLSOptions: nil,
	},
	{
		Address: "0.0.0.0:8883",
	},
}

// LogConfig is use to configure the log behaviors.
type LogConfig struct {
	// Level is the log level. Possible values: debug, info, warn, error
	Level string `yaml:"level"`
	// Format is the log format. Possible values: json, text
	Format string `yaml:"format"`
	// DumpPacket indicates whether to dump MQTT packet in debug level.
	DumpPacket bool `yaml:"dump_packet"`
}

func (l LogConfig) Validate() error {
	if l.Level != "debug" && l.Level != "info" && l.Level != "warn" && l.Level != "error" {
		return fmt.Errorf("invalid log level: %s", l.Level)
	}
	if l.Format != "json" && l.Format != "text" {
		return fmt.Errorf("invalid log format: %s", l.Format)
	}
	return nil
}

// Config is the configration for gmqttd.
type Config struct {
	Listeners         []*ListenerConfig `yaml:"listeners"`
	MQTT              MQTT              `yaml:"mqtt,omitempty"`
	Log               LogConfig         `yaml:"log"`
	ConfigDir         string            `yaml:"config_dir"`
	Persistence       Persistence       `yaml:"persistence"`
	TopicAliasManager TopicAliasManager `yaml:"topic_alias_manager"`
	DumpPacket        bool              `yaml:"dump_packet"`
}

type GRPC struct {
	Endpoint string `yaml:"endpoint"`
}

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

type ListenerConfig struct {
	Address     string `yaml:"address"`
	*TLSOptions `yaml:"tls"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type config Config
	raw := config(DefaultConfig())
	if err := unmarshal(&raw); err != nil {
		return err
	}
	emptyMQTT := MQTT{}
	if raw.MQTT == emptyMQTT {
		raw.MQTT = DefaultMQTTConfig
	}
	*c = Config(raw)
	return nil
}

func (c Config) Validate() (err error) {
	err = c.Log.Validate()
	if err != nil {
		return err
	}
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

func ParseConfig(filePath string) (c Config, err error) {
	if filePath == "" {
		return DefaultConfig(), nil
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return c, err
	}
	c = DefaultConfig()
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return c, err
	}
	c.ConfigDir = path.Dir(filePath)
	err = c.Validate()
	if err != nil {
		return Config{}, err
	}
	return c, err
}
