package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var cfg *Config = &Config{}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("init config err, can not get pwd")
		return
	}

	cfg.PWD = pwd
}

func LoadYAMLConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, cfg)
	return err
}

func GetConfig() *Config {
	return cfg
}
