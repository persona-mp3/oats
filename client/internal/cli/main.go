package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/persona-mp3/client/shared"
)

func defaultCredentials() *shared.Credentials {
	return &shared.Credentials{
		Username: "master_user",
		Password: "m@ster_password",
	}
}

func ParseLoginCredentials() *shared.Credentials {
	var userName string
	var password string
	creds := &shared.Credentials{}

	fmt.Println("provide credentials")
	fmt.Printf("username: ")
	fmt.Scanf("%s", &userName)

	fmt.Printf("password: ")
	fmt.Scanf("%s", &password)

	_username := strings.ReplaceAll(userName, "", "")
	_password := strings.ReplaceAll(password, "", "")

	if len(_username) < 3 || len(_password) < 3 {
		fmt.Println("using default credentials")
		return defaultCredentials()
	}

	creds.Username = userName
	creds.Password = password

	return creds
}

func ReadArgs() (string, bool) {
	var endpoint string
	var config bool
	flag.StringVar(&endpoint, "ep", shared.WelcomeEndpoint, "endpoint to visit, default is welcome")
	flag.BoolVar(&config, "cfg", false, "load credentials from .oat_creds.json")
	flag.Parse()

	if len(endpoint) < 4 {
		fmt.Println("using default endpoint")
		endpoint = shared.WelcomeEndpoint
	}

	return endpoint, config
}

// Loads credentials in the current directory under the name of .oat_creds.json
func LoadCredentials() (*shared.Credentials, error) {
	// open the file with credentials
	content, err := os.ReadFile(".oat_creds.json")
	if err != nil {
		return nil, fmt.Errorf(" could not open .oat_creds.json: %w", err)
	}

	creds := &shared.Credentials{}
	if err := json.Unmarshal(content, creds); err != nil {
		return nil, fmt.Errorf(" could not parse config: %w", err)
	}

	fmt.Println(" credentials loaded for: ", creds.Username, creds.Password)
	return creds, nil

}
