package handlers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Prompt   string `yaml:"prompt"`
	Default  string `yaml:"default"`
	Commands []struct {
		Command   string `yaml:"command"`
		Response  string `yaml:"response"`
		Color     string `yaml:"color"`
		Delay     int    `yaml:"delay"`
		Responses []struct {
			Response string `yaml:"response"`
			Color    string `yaml:"color"`
			Delay    int    `yaml:"delay"`
		} `yaml:"responses"`
	} `yaml:"commands"`
}

func LoadConfig(config *Config) {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		readConfig(fileName, config)
		go watchConfig(fileName, config)
	} else {
		fmt.Println("Error: no protocli config given")
		fmt.Print("\nprotocli <config>\n\n")
		os.Exit(1)
	}
}

func readConfig(fileName string, config *Config) {

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

}

func watchConfig(fileName string, config *Config) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Error: file watcher %s", err.Error())
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					readConfig(fileName, config)
				}
			case err := <-watcher.Errors:
				fmt.Printf("Error: file watcher event %s", err.Error())
				os.Exit(1)
			}

		}

	}()

	err = watcher.Add(fileName)
	if err != nil {
		fmt.Printf("Error: file watcher add %s", err.Error())
		os.Exit(1)
	}
	<-done

}
