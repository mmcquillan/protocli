package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mmcquillan/matcher"
	"github.com/mmcquillan/protocli/handlers"
)

var version string

func main() {

	// strarting
	fmt.Println("protocli " + version)
	if len(os.Args) == 2 && os.Args[1] == "version" {
		os.Exit(0)
	}

	// handler
	handlers.Exit()
	config := handlers.LoadConfig()

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
					if cmd.Response != "" {
						fmt.Println(handlers.Substitute(cmd.Response, values))
					}
					if len(cmd.Responses) > 0 {
						for _, r := range cmd.Responses {
							fmt.Println(handlers.Substitute(r.Response, values))
						}
					}
				}
			}

		}
		fmt.Print("\n" + config.Prompt)
	}
}
