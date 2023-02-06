package cfg

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	filename = "/config.yaml"
)

var Config config

type serverCfg struct {
	Host      string `yaml:"host"`
	ForceStop int    `yaml:"force_stop"`
	Grpc      struct {
		Port int `yaml:"port"`
	} `yaml:"grpc"`

	CaptchaServer struct {
		Name    string `yaml:"name"`
		LogPath string `yaml:"log_path"`

		Expires int    `yaml:"expires"`
		Port    int    `yaml:"port"`
		Subject string `yaml:"subject"`
		Format  string `yaml:"format"`
		Host    string `yaml:"host"`
		From    string `yaml:"from"`
		Secret  string `yaml:"secret"`
	} `yaml:"captcha_server"`

	AccountServer struct {
		Name    string `yaml:"name"`
		LogPath string `yaml:"log_path"`
	} `yaml:"account_server"`
}

type config struct {
	Server serverCfg `yaml:"server"`

	Log struct {
		Level  int  `yaml:"level"`
		Color  bool `yaml:"color"`
		Caller bool `yaml:"caller"`
	} `yaml:"log"`

	Etcd struct {
		Endpoints []string `yaml:"endpoints"`
		Timeout   int
	} `yaml:"etcd"`

	Redis struct {
		Addr string `yaml:"addr"`
		Pass string `yaml:"pass"`
		DB   int    `yaml:"db"`
	} `yaml:"redis"`

	Mysql struct {
		Port   int    `yaml:"port"`
		Host   string `yaml:"host"`
		Name   string `yaml:"name"`
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
	} `yaml:"mysql"`
}

func InitConfig() {
	var path, defaultPath string
	flag.StringVar(&path, "cfg_path", "./cfg", "log path")
	flag.StringVar(&defaultPath, "default_path", ".config", "default log path")
	flag.Parse()
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Config); err != nil {
		panic(err)
	}

	initDefaultConfig(defaultPath)
}

func (cfg config) Dsn() string {
	db := cfg.Mysql
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Host, db.Port, db.Name) +
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
}

func initDefaultConfig(defaultPath string) {
	data, err := ioutil.ReadFile(defaultPath)
	if err != nil {
		return
	}
	_ = yaml.Unmarshal(data, &Config)
}
