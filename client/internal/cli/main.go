package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/persona-mp3/client/common"
)

const (
// welcomeEndpoint = "/welcom"
)

func defaultCredentials() *common.Credentials {
	return &common.Credentials{
		Username: "master_user",
		Password: "m@ster_password",
	}
}

func ParseLoginCredentials() *common.Credentials {
	var userName string
	var password string
	creds := &common.Credentials{}

	fmt.Println(" provide credentials")
	for {
		fmt.Printf(" username: ")
		fmt.Scanf("%s", &userName)

		fmt.Printf(" password: ")
		fmt.Scanf("%s", &password)

		_username := strings.ReplaceAll(userName, " ", "")
		_password := strings.ReplaceAll(password, " ", "")

		if len(_username) < 3 || len(_password) < 3 {
			fmt.Println(" using default credentials")
			return defaultCredentials()
		}

		break
	}
	creds.Username = userName
	creds.Password = password

	return creds
}

func ReadArgs() (string, bool) {
	var endpoint string
	var config bool
	flag.StringVar(&endpoint, "ep", "/welcome", "endpoint to visit, default is welcome")
	flag.BoolVar(&config, "cfg", false, " load credentials from .oat_creds.json")
	flag.Parse()

	// TODO: validate endpoint
	if len(endpoint) < 4 {
		fmt.Println(" using default endpoint")
		endpoint = common.WelcomeEndpoint
	}


	return endpoint, config
}

