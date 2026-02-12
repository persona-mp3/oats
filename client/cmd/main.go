package main

import (
	"fmt"
	"log"

	"github.com/persona-mp3/client/internal/api"
	"github.com/persona-mp3/client/internal/cli"
	"github.com/persona-mp3/client/shared"
	"github.com/persona-mp3/client/websocket"
)

func main() {
	endpoint, loadCreds := cli.ReadArgs()

	if endpoint != "/login" {
		fmt.Printf("%s is still under construction\n", endpoint)
		return
	}

	creds := &shared.Credentials{}
	var err error

	if !loadCreds {
		creds = cli.ParseLoginCredentials()
	} else {
		creds, err = cli.LoadCredentials()
	}

	if err != nil {
		log.Fatal(err)
	}

	info, err := api.LoginHandler(creds)
	if err != nil {
		log.Fatal(err)
	}

	if info == nil {
		log.Fatalf(" yo bro, %+v\n", info)
	}
	if err := websocket.StartWebSocketProtocol(info); err != nil {
		log.Fatal(err)
	}

}
