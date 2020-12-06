package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var cfg *Config = &Config{}

//FIXME: cannot read yaml
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
