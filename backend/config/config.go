package config

import (
	"backend/consts"
	"backend/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type YamlConfig struct {
	Database struct {
		Url      string `yaml:"url"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"db_name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var yamlConfig YamlConfig

func Init(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return err
	}

	log.Info(consts.Config, "Read config.yaml successfully")
	return nil
}

func GetYamlConfig() YamlConfig {
	return yamlConfig
}

func Dsn(y YamlConfig) string {
	return y.Database.Username + ":" + y.Database.Password + "@tcp(" + y.Database.Url + ":" + y.Database.Port + ")/" + y.Database.DbName
}
