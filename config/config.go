package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Conf struct {
	Port  int32  `yaml:"port"`
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
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
			Port:  9000,
			Host:  "127.0.0.1",
			Token: "",
		})
		ioutil.WriteFile(
			"config.yml",
			defaultConfig,
			777,
		)
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
