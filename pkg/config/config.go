package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ServerAddress    string   `yaml:"server_address"`
	LoginGateAddress string   `yaml:"login_gate_address"`
	DatabaseURL      string   `yaml:"database_url"`
	RedisURL         string   `yaml:"redis_url"`
	KafkaURL         []string `yaml:"kafka_url"`
	ClusterNodes     []string `yaml:"cluster_nodes"`
	WorkerPoolSize   int      `yaml:"worker_pool_size"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
