package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func GenerateConfig() {
	config := Config{
		Username:    "",
		AccessToken: "",
	}
	jsonConfig, err := json.MarshalIndent(config, "", "")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./config.json", jsonConfig, 0777)
	if err != nil {
		panic(err)
	}
}

func ReadConfig() Config {
	config := Config{}
	fileExist := IsFileExist("config.json")
	if !fileExist {
		GenerateConfig()
	}
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &config)
	return config
}
