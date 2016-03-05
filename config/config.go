package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	DatabaseIP       string `json:"databaseIP"`
	DatabaseUser     string `json:"databaseUser"`
	DatabasePassword string `json:"databasePassword"`
	DatabaseName     string `json:"databaseName"`
	IsProduction     bool   `json:"isProduction"`
}

var ConfigData Config

func ParseConfig() Config {
	file, err := ioutil.ReadFile("/Users/wkrause/Documents/go/src/apretaste/conf.json")
	if err != nil {
		log.Fatal("Could not find config file")
	}
	var data Config
	json.Unmarshal(file, &data)
	return data
}
