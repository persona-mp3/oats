package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/persona-mp3/client/shared"
)

const (
	authAddr  = "http://localhost:8000"
	loginPath = "/login"
)

func parseUrl(addr, endpoint string) (string, error) {
	_url, err := url.Parse(addr)
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	fullEndpoint := _url.JoinPath(endpoint)
	return fullEndpoint.String(), nil
}

func getRedirectUrl(client *http.Client) *shared.RedirectInfo {
	info := &shared.RedirectInfo{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		to, err := req.Response.Location()
		if err != nil {
			return err
		}

		info.Url = to
		info.StatusCode = req.Response.StatusCode

		return http.ErrUseLastResponse
	}
	return info
}

// Sends the credentials in a post request, using a http Client
// to the Auth server for verification. Upon verification,
// the client is redirected to the WebSocket server, but
// this url is stored in RedirectInfo so that a proper
// WebSocket Client can be made to perform the request to the server
func contactLoginEndpoint(creds *shared.Credentials) (*shared.RedirectInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	endpoint, err := parseUrl(authAddr, loginPath)
	if err != nil {
		return nil, err
	}

	fmt.Println("contacting ", endpoint)

	payload, err := json.Marshal(&creds)
	if err != nil {
		return nil, fmt.Errorf("could not marshall credentials %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("content-type", "application/json")
	client := &http.Client{}

	redirectInfo := getRedirectUrl(client)

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occured in contacting server: %w", err)
	}

	if response.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error from server %d", response.StatusCode)
	}

	return redirectInfo, nil
}

func LoginHandler(creds *shared.Credentials) (*shared.RedirectInfo, error) {
	info, err := contactLoginEndpoint(creds)
	if err != nil {
		return nil, err
	}
	return info, nil
}
