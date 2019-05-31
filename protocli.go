package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
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
	var config handlers.Config
	handlers.LoadConfig(&config)

	// initialize cli
	fmt.Print("\n" + config.Prompt)
	scanner := bufio.NewScanner(os.Stdin)

	// input loop
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		// help
		if input == "?" {
			for _, cmd := range config.Commands {
				fmt.Println(cmd.Command)
			}

			// find match
		} else if input != "" {
			cmdMatch := false
			for _, cmd := range config.Commands {
				if !cmdMatch {

					match, _, values := matcher.Matcher(cmd.Command, input)
					if match {

						cmdMatch = true

						// delay
						time.Sleep(time.Second * time.Duration(cmd.Delay))

						// sub response
						response := handlers.Substitute(cmd.Response, values)

						// single response
						if cmd.Response != "" {
							switch strings.ToLower(cmd.Color) {
							case "red":
								color.Red(response)
							case "green":
								color.Green(response)
							case "yellow":
								color.Yellow(response)
							case "blue":
								color.Blue(response)
							case "magenta":
								color.Magenta(response)
							case "cyan":
								color.Cyan(response)
							case "white":
								color.White(response)
							default:
								fmt.Println(response)
							}
						}

						// list response
						if len(cmd.Responses) > 0 {
							for _, r := range cmd.Responses {

								// delay
								time.Sleep(time.Second * time.Duration(r.Delay))

								// sub response
								response := handlers.Substitute(r.Response, values)

								// single response
								switch strings.ToLower(r.Color) {
								case "red":
									color.Red(response)
								case "green":
									color.Green(response)
								case "yellow":
									color.Yellow(response)
								case "blue":
									color.Blue(response)
								case "magenta":
									color.Magenta(response)
								case "cyan":
									color.Cyan(response)
								case "white":
									color.White(response)
								default:
									fmt.Println(response)
								}

							}
						}
					}
				}

				// print default if no matches
				if !cmdMatch && config.Default != "" {
					fmt.Println(config.Default)
				}

			}
		}
		fmt.Print("\n" + config.Prompt)
	}
}
