package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Conf struct {
	Host        string `yaml:"host"`
	HttpApiPort int32  `yaml:"http_api_port"`
	WsPort      int32  `yaml:"ws_port"`
	Token       string `yaml:"token"`
}

var Config *Conf = nil

func LoadConf() (*Conf, error) {
	if Config == nil {
		Config = &Conf{}
		return Config.loadConf()
	} else {
		return Config, nil
	}
}

func (c *Conf) loadConf() (*Conf, error) {
	filename := "config.yml"

	if !checkFileIsExist(filename) {
		defaultConfig, _ := yaml.Marshal(Conf{
			Host:        "127.0.0.1",
			WsPort:      9000,
			HttpApiPort: 9001,
			Token:       "",
		})
		ioutil.WriteFile(
			"config.yml",
			defaultConfig,
			777,
		)
		return nil, errors.New("'config.yml' not found! Created default config. Edit it please")
	}
	confFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(confFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
