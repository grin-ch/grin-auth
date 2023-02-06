package cfg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const (
	grin_config_path         = "GRIN_CFG_PATH"
	grin_default_config_path = "GRIN_DEFAULT_CFG_PATH"
	filename                 = "config.yaml"
)

var Config config

type serverInfo struct {
	Name     string `yaml:"name"`
	GrpcPort int    `yaml:"grpc_port"`
	LogPath  string `yaml:"log_path"`
}

type serverCfg struct {
	Host      string `yaml:"host"`
	ForceStop int    `yaml:"force_stop"`

	CaptchaServer struct {
		Info serverInfo `yaml:"info"`

		Expires int    `yaml:"expires"`
		Port    int    `yaml:"port"`
		Subject string `yaml:"subject"`
		Format  string `yaml:"format"`
		Host    string `yaml:"host"`
		From    string `yaml:"from"`
		Secret  string `yaml:"secret"`
	} `yaml:"captcha_server"`

	AccountServer struct {
		Info serverInfo `yaml:"info"`

		Expires int    `yaml:"expires"`
		Signed  string `yaml:"signed"`
		Issuer  string `yaml:"issuer"`
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

func (cfg config) Dsn() string {
	db := cfg.Mysql
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Host, db.Port, db.Name) +
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
}

func InitConfig() {
	envPath := os.Getenv(grin_config_path)
	defaultPath := os.Getenv(grin_default_config_path)
	data, err := ioutil.ReadFile(path.Join(envPath, filename))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Config); err != nil {
		panic(err)
	}

	initDefaultConfig(defaultPath)
}

func initDefaultConfig(defaultPath string) {
	data, err := ioutil.ReadFile(path.Join(defaultPath, filename))
	if err != nil {
		return
	}
	_ = yaml.Unmarshal(data, &Config)
}
