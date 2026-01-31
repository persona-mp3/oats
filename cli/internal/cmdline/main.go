package cmdline

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/persona-mp3/cli/internal/api"
	"github.com/persona-mp3/cli/internal/utils"
)

const (
	SERVER_ADDR   = "http://localhost:8900"
	loginRoute    = "/login"
	registerRoute = "/register"
)

func ParseArgs() string {
	var endpoint string
	var defaultCreds string
	flag.StringVar(&endpoint, "ep", "/welcome", "endpoint to visit, default is /welcome")
	flag.StringVar(&defaultCreds, "def", "", "uses default credentials")

	// TODO: check against whitespaces
	if len(endpoint) <= 5 {
		fmt.Println("using default endpoint")
		endpoint = "/welcome"
	}
	return endpoint
}

var masterCreds = api.UserCredentials{
	UserName: "childish_gambino",
	Password: "awaken_my_love",
	Email:    "donaldglover@spotify.com",
}

// this is called when the tgtEndpoint is post request for /login & /register
func parseCredentials() *api.UserCredentials {
	var userName string
	var email string
	var passwd string
	var defaultCreds string

	credentials := &api.UserCredentials{}

	fmt.Println(" please provide credentials to authenticate")
	fmt.Println(" use master credentials [y/n]?:")

	for {
		fmt.Scanf("%s", defaultCreds)

		if defaultCreds == "y" {
			fmt.Println(" using master credentials")
			return &masterCreds
		}

		fmt.Println(" username: ")
		fmt.Scanf("%s", strings.ReplaceAll(userName, " ", ""))

		fmt.Println(" password: ")
		fmt.Scanf("%s", strings.ReplaceAll(passwd, " ", ""))

		fmt.Println(" email: ")
		fmt.Scanf("%s", strings.ReplaceAll(email, " ", ""))

		// cheesy validations for now
		if len(userName) <= 1 || len(passwd) <= 1 ||
			len(email) <= 1 || !strings.Contains(email, "@") {
			fmt.Println("bro...cmon, provide valid credentials")
			continue
		}

		credentials.UserName = userName
		credentials.Password = passwd
		credentials.Email = email
		break

	}
	return credentials
}

// Depending on the endpoint the user wants to visit
// there would most definitely be different conditions
// but for now, it's mostly the /login && /register route
//
//	func MapperFunc(endpoint string) string {
//		tgtEndpoint := ParseArgs()
//
//		if tgtEndpoint == loginRoute {
//			parseCredentials()
//		} else
//
//		return ""
//	}
func HandleLoginRoute(endpoint string) (*api.Req, error) {
	creds := parseCredentials()
	data, err := utils.ToJson(&creds)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	req := &api.Req{
		Method:   http.MethodPost,
		Endpoint: endpoint,
		Body:     reader,
	}

	return req, nil
}

func MapTo(endpoint string) string {
	if endpoint != loginRoute && endpoint != registerRoute {
		return "general"
	}

	return endpoint
}
