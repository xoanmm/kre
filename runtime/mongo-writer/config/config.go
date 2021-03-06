package config

import (
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var once sync.Once
var cfg *Config

// Config holds the configuration values for the application
type Config struct {
	Debug string `yaml:"debug" envconfig:"DEBUG"`
	Nats  struct {
		Server    string `yaml:"server" envconfig:"KRT_NATS_SERVER"`
		ClusterID string `yaml:"clusterId" envconfig:"KRT_NATS_CLUSTER_ID"`
	} `yaml:"nats"`

	MongoDB struct {
		Address     string `yaml:"address" envconfig:"KRE_RUNTIME_MONGO_URI"`
		DBName      string `yaml:"dbName" envconfig:"KRE_MONGODB_DB_NAME"`
		ConnTimeout int    `yaml:"connTimeout" envconfig:"KRE_MONGODB_CONN_TIMEOUT"`
	} `yaml:"mongodb"`
}

// NewConfig will read the config.yml file and override values with env vars.
func NewConfig() *Config {
	once.Do(func() {
		f, err := os.Open("config.yml")
		if err != nil {
			panic(err)
		}

		cfg = &Config{}
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(cfg)
		if err != nil {
			panic(err)
		}

		err = envconfig.Process("", cfg)
		if err != nil {
			panic(err)
		}
	})

	return cfg
}
