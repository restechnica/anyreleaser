package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(path string) (config Root, err error) {
	return Overload(path, config)
}

func Overload(path string, config Root) (overloaded Root, err error) {
	var data []byte

	if data, err = ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(data, &config)
	}

	return config, err
}
