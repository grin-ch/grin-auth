package cfg

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	configPath        = "config.yaml"
	defaultConfigPath = "../.config/config.yaml"
)

var Config config

type config struct {
	Server struct {
		Host      string `yaml:"host"`
		ForceStop int    `yaml:"force_stop"`
		Grpc      struct {
			Port int `yaml:"port"`
		} `yaml:"grpc"`
	} `yaml:"server"`

	Etcd struct {
		Endpoints []string `yaml:"endpoints"`
		Timeout   int
	} `yaml:"etcd"`
}

func InitConfig() {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Config); err != nil {
		panic(err)
	}

	initDefaultConfig()
}

func initDefaultConfig() {
	data, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		return
	}
	_ = yaml.Unmarshal(data, &Config)
}
