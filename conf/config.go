package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	ServerConfPath = "./conf/server.yaml"
	LogConfPath    = "./conf/log.yaml"
)

type (
	ServerConf struct {
		Port         string `yaml:"port"`
		ShutDownWait int64  `yaml:"shutdown_wait"`
	}

	LogConf struct {
		Zap []ZapLogConf `yaml:"zap_log"`
	}

	ZapLogConf struct {
		Level      string `yaml:"level"`
		Filename   string `yaml:"file_name"`
		MaxSize    int    `yaml:"max_size"`
		MaxAge     int    `yaml:"max_age"`
		MaxBackups int    `yaml:"max_backups"`
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
