package main

import (
	"chatserver/client"
	"chatserver/tui"
	"flag"
	"log"
)

func main() {
	address := flag.String("server", "", "Which server to connect to")
	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	tui.StartUI(client)
}
