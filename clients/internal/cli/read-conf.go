package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/persona-mp3/client/common"
)

func LoadCredentials() (*common.Credentials, error) {
	// open the file with credentials
	content, err := os.ReadFile(".oat_creds.json")
	if err != nil {
		return nil, fmt.Errorf(" could not open .oat_creds.json: %w", err)
	}

	creds := &common.Credentials{}
	if err := json.Unmarshal(content, creds); err != nil {
		return nil, fmt.Errorf(" could not parse config: %w", err)
	}

	fmt.Println(" credentials loaded for: ", creds.Username, creds.Password)
	return creds, nil

}
