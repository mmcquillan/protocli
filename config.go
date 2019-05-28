package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Prompt   string `yaml:"prompt"`
	Commands []struct {
		Command   string `yaml:"command"`
		Response  string `yaml:"response"`
		Responses []struct {
			Response string `yaml:"response"`
		} `yaml:"responses"`
	} `yaml:"commands"`
}

func GetConfig(config *Config) {

	// load config
	file, err := ioutil.ReadFile("./protocli.yaml")
	if err != nil {
		fmt.Printf("config file read: %s", err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("config file yaml unmarshal: %s", err.Error())
		os.Exit(1)
	}

	// set defaults
	if config.Prompt == "" {
		config.Prompt = "> "
	}

}
