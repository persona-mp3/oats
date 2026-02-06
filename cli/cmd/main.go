package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/persona-mp3/cli/internal/api"
	"github.com/persona-mp3/cli/internal/utils"
)

const (
	POST_METHOD   = "POST"
	GET_METHOD    = "GET"
	addr          = "http://localhost:8990"
	docsRoute     = "/docs"
	helpRoute     = "/help-me"
	loginRoute    = "/login"
	registerRoute = "/register"
	welcomeRoute  = "/welcome"
)

func newClient() *http.Client {
	return &http.Client{}
}

func parseCmdLine() string {
	// later on, add validation to check if the endpoint actually exists
	var ep string
	flag.StringVar(&ep, "ep", "/welcome", "endpoint to visit, default is welcome")
	flag.Parse()

	if len(ep) < 5 {
		fmt.Println("contacting default route")
		ep = welcomeRoute
	}

	return ep
}

func _masterCredentials() *api.UserCredentials {
	return &api.UserCredentials{
		UserName: "childish_gambino",
		Password: "awaken_my_love",
		Email:    "donaldglover@spotify.com",
	}
}

func parseLoginCredentials() *api.UserCredentials {
	var userName string
	var passwd string
	var email string

	for {
		fmt.Print("\n provide userName: ")
		fmt.Scanf("%s", &userName)

		fmt.Print("\n password: ")
		fmt.Scanf("%s", &passwd)

		fmt.Print("\n email: ")
		fmt.Scanf("%s", &email)

		if len(userName) < 3 || len(passwd) < 3 || len(email) < 3 {
			fmt.Println(" using _masterCredentials")
			return _masterCredentials()
		}

		break
	}

	return &api.UserCredentials{
		UserName: userName,
		Password: passwd,
		Email:    email,
	}
}

func main() {
	tgtEndpoint := parseCmdLine()

	newReq := &api.Req{
		Method:   http.MethodGet,
		Endpoint: addr + tgtEndpoint,
		Body:     nil,
	}
	if tgtEndpoint == registerRoute || tgtEndpoint == loginRoute {
		fmt.Printf(" hitting %s, provide credentials\n", tgtEndpoint)
		creds := parseLoginCredentials()
		data, err := utils.ToJson(creds)

		if err != nil {
			log.Fatal(err)
		}

		newReq.Method = http.MethodPost
		newReq.Body = bytes.NewReader(data)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newReq.Context = ctx

	client := newClient()

	if tgtEndpoint == registerRoute {
		if err := newReq.HandleLoginRoute(client); err != nil {
			log.Fatal(err)
		}
		return

	} else if tgtEndpoint == loginRoute {
		fmt.Println(" working on login route")
		return
	}

	res, err := newReq.GenericRequest(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" response from server", res.StatusCode)
	fmt.Printf("\n %s\n", res.Content)
}
