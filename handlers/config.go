package handlers

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
		Color     string `yaml:"color"`
		Responses []struct {
			Response string `yaml:"response"`
			Color    string `yaml:"color"`
		} `yaml:"responses"`
	} `yaml:"commands"`
}

func LoadConfig() (config *Config) {

	if len(os.Args) > 1 {
		fileName := os.Args[1]

		// load config
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error: config file read: %s", err.Error())
			os.Exit(1)
		}
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			fmt.Printf("Error: config file yaml unmarshal: %s", err.Error())
			os.Exit(1)
		}

		// set defaults
		if config.Prompt == "" {
			config.Prompt = "> "
		}

	} else {
		fmt.Println("Error: no protocli config given")
		fmt.Print("\nprotocli <config>\n\n")
		os.Exit(1)
	}

	return config
}
