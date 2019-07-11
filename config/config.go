package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port     string   `json:"port"`
	DBConfig DBConfig `json:"dbConfig"`
}

func ReadConfigFile(filename string, out interface{}) error {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(configData, out)
}

type DBConfig struct {
	Driver   string `json:"dbDriver"`
	Host     string `json:"dbHost"`
	Port     string `json:"dbPort"`
	User     string `json:"dbUser"`
	Password string `json:"dbPassword"`
	Name     string `json:"dbName"`
}
