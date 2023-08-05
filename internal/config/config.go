package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type RabbitMQConfig struct {
	URl      string      `yaml:"url"`
	QueueCfg QueueConfig `yaml:"queue_config"`
}

type QueueConfig struct {
	Exchange  string `yaml:"exchange"`
	QueueName string `yaml:"queue_name"`
}

type ServiceConfig struct {
	Port uint64 `yaml:"port"`
}

type Summary struct {
	RabbitMQConfig `yaml:"rabbit_config"`
	ServiceConfig  `yaml:"service_config"`
}

func ParseConfig() (*Summary, error) {
	filename, err := filepath.Abs(".swampy/values.yml")

	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Summary{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	logrus.Info("config successfully parsed")
	return config, nil
}
