package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/mmcquillan/matcher"
	"gopkg.in/yaml.v2"
)

func main() {

	// strarting
	fmt.Println("protocli")

	// load config
	var config Config
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

	// loop on input
	fmt.Print("\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input != "" {
			for _, cmd := range config.Commands {
				match, _, values := matcher.Matcher(cmd.Command, input)
				if match {
					fmt.Println(substitute(cmd.Response, values))
				}
			}

		}
		fmt.Print("\n> ")
	}
}

type Config struct {
	Commands []struct {
		Command  string `yaml:"command"`
		Response string `yaml:"response"`
	} `yaml:"commands"`
}

func substitute(value string, tokens map[string]string) string {
	if match, hits := findVars(value); match {
		for _, hit := range hits {
			value = strings.Replace(value, hit, tokens[strip(hit)], -1)
		}
	}
	return value
}

func findVars(value string) (match bool, tokens []string) {
	match = false
	re := regexp.MustCompile("\\${([A-Za-z0-9:*_\\|\\-\\.\\?]+)}")
	tokens = re.FindAllString(strings.Replace(value, "$${", "X{", -1), -1)
	if len(tokens) > 0 {
		match = true
	}
	return match, tokens
}

func strip(value string) (stripped string) {
	stripped = strings.Replace(value, "${", "", -1)
	stripped = strings.Replace(stripped, "}", "", -1)
	return stripped
}
