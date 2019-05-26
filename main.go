package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mmcquillan/matcher"
)

func main() {

	// strarting
	fmt.Println("protocli")

	// load config
	var config Config
	GetConfig(&config)

	// loop on input
	fmt.Print("\n" + config.Prompt)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "?" {
			for _, cmd := range config.Commands {
				fmt.Println(cmd.Command)
			}
		} else if input != "" {
			for _, cmd := range config.Commands {
				match, _, values := matcher.Matcher(cmd.Command, input)
				if match {
					fmt.Println(Substitute(cmd.Response, values))
				}
			}

		}
		fmt.Print("\n" + config.Prompt)
	}
}
