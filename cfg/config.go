package cfg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	grin_config_path         = "GRIN_CFG_PATH"
	grin_default_config_path = "GRIN_DEFAULT_CFG_PATH"
	filename                 = "config.yaml"
)

var Config config

type ServerInfo struct {
	Name        string `yaml:"name"`
	Port        int    `yaml:"port"`
	LogPath     string `yaml:"log_path"`
	PprofEnable bool   `yaml:"pprof_enable"`
	PprofPort   int    `yaml:"pprof_port"`
}

type serverCfg struct {
	Host      string `yaml:"host"`
	ForceStop int    `yaml:"force_stop"`

	HttpServer struct {
		Info ServerInfo `yaml:"info"`

		Mode           string         `yaml:"mode"`
		Debug          bool           `yaml:"debug"`
		Timeout        int            `yaml:"timeout"`
		TimeoutAppoint map[string]int `yaml:"timeout_appoint"`
	} `yaml:"http_server"`

	CaptchaServer struct {
		Info ServerInfo `yaml:"info"`

		Expires int    `yaml:"expires"`
		Port    int    `yaml:"port"`
		Subject string `yaml:"subject"`
		Format  string `yaml:"format"`
		Host    string `yaml:"host"`
		From    string `yaml:"from"`
		Secret  string `yaml:"secret"`
	} `yaml:"captcha_server"`

	AccountServer struct {
		Info ServerInfo `yaml:"info"`

		ForceCheck bool   `yaml:"force_check"` // 强制登入校验验证码
		Expires    int    `yaml:"expires"`
		Signed     string `yaml:"signed"`
		Issuer     string `yaml:"issuer"`
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
	rand.Seed(time.Now().UnixNano())

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
