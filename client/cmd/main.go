package main

import (
	"fmt"
	"log"

	"github.com/persona-mp3/client/common"
	"github.com/persona-mp3/client/internal/api"
	"github.com/persona-mp3/client/internal/cli"
	"github.com/persona-mp3/client/oats"
)

func main() {
	tgtEndpoint := cli.ReadArgs()

	if tgtEndpoint != common.LoginEndpoint {
		fmt.Println("handling for other route")
		return
	}
	creds := cli.ParseLoginCredentials()
	wsAddr, err := api.HandleLoginRoute(creds)
	if err != nil {
		log.Fatal(err)
	}

	if err := oats.BeginOatsProtocol(wsAddr); err != nil {
		log.Fatal(err)
	}
}
