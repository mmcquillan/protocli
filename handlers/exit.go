package handlers

import (
	"fmt"
	"os"
	"os/signal"
)

func Exit() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	go func(signalChannel chan os.Signal) {
		sig := <-signalChannel
		fmt.Println("\nprotocli exiting ... [" + sig.String() + "]")
		os.Exit(0)
	}(signalChannel)

}
