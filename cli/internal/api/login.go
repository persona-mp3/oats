package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func configureRedirectToWSS(c *http.Client) *RedirectInfo {
	info := &RedirectInfo{}

	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectUrl, err := req.Response.Location()
		if err != nil {
			return fmt.Errorf("could not get redirect-url\n: %w", err)
		}

		info.url = redirectUrl
		info.statusCode = req.Response.StatusCode

		return http.ErrUseLastResponse
	}
	return info
}

func (req *Req) LoginRouteHandler(c *http.Client) (*GenericRes, *RedirectInfo, error) {
	fmt.Println("\n making request to:", req.Endpoint)
	if c == nil {
		return nil, nil, fmt.Errorf("client pointer is nil")
	}

	info := configureRedirectToWSS(c)

	newReq, err := http.NewRequestWithContext(req.Context, req.Method, req.Endpoint, req.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error in creating new request:\n: %w", err)
	}

	newReq.Header.Set("Content-Type", "application/json")

	response, err := c.Do(newReq)
	if err != nil {
		return nil, nil, fmt.Errorf("error occured contacting server\n: %w", err)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error occured reading response from server\n: %w", err)
	}

	serverRes := &GenericRes{}

	serverRes.StatusCode = response.StatusCode
	serverRes.Content = content

	return serverRes, info, nil
}

func (req *Req) HandleLoginRoute(c *http.Client) error {
	if c == nil {
		log.Fatalf("\n [panic] client pointer is null\n")
	}
	res, redirectInfo, err := req.LoginRouteHandler(c)
	if err != nil {
		return err
	}

	log.Printf("status-code: %d\n", res.StatusCode)
	log.Printf("server response: %s\n", res.Content)

	log.Printf("handling upgrade...\n")
	if redirectInfo == nil {
		log.Fatal("we've caught the culprit, they said")
	}

	// Making a websocket request
	_err := req.requestUpgrade(redirectInfo)
	if _err != nil {
		return err
	}
	return nil
}
