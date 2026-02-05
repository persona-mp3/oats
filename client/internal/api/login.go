package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/persona-mp3/client/common"
)

func configRedirectToWSS(c *http.Client) *common.RedirectInfo {
	info := &common.RedirectInfo{}

	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		to, err := req.Response.Location()
		if err != nil {
			return err
		}

		info.Location = to
		return http.ErrUseLastResponse
	}

	return info
}

func ToJson(value any) ([]byte, error) {
	content, err := json.Marshal(&value)
	if err != nil {
		return []byte{}, fmt.Errorf(" could not marhsal: \n%w", err)
	}
	return content, nil
}

func createRequest(creds *common.Credentials) (*common.RedirectInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// TODO: parseUrl function
	loginRoute := common.AuthServer + common.LoginEndpoint

	content, err := ToJson(creds)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginRoute, bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf(" could not create request:\n  %w", err)
	}

	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	info := configRedirectToWSS(client)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(" error in executing request:\n  %w", err)
	}

	if res.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf(" error from server: %d\n", res.StatusCode)
	}

	defer res.Body.Close()

	fmt.Printf(" redirect-info: %+v\n", info)

	return info, nil
}

func HandleLoginRoute(creds *common.Credentials) (string, error) {
	info, err := createRequest(creds)

	if err != nil {
		return "", err
	}

	return info.Location.String(), nil
}
