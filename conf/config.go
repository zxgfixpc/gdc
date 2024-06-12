package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	ServerPath = "./conf/server.yaml"
)

type (
	Config struct {
	}

	ServerConf struct {
		Port         string `yaml:"port"`
		ShutDownWait int64  `yaml:"shutdown_wait"`
	}
)

func Parser(conf interface{}, path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("yaml file get err %v", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return fmt.Errorf("yaml unmarshal err %v", err)
	}

	return nil
}
