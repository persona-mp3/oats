package api

import (
	"fmt"
	"io"
	"net/http"
)

func configureRedirectToWSS(c *http.Client) (*RedirectInfo, error) {
	info := &RedirectInfo{}

	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectUrl, err := req.Response.Location()
		if err != nil {
			return fmt.Errorf("error in getting redirect-url\n: %w", err)
		}

		info.url = redirectUrl
		info.statusCode = req.Response.StatusCode
		return http.ErrUseLastResponse
	}
	return info, nil
}

func (req *Req) LoginRouteHandler(c *http.Client) (*GenericRes, *RedirectInfo, error) {
	fmt.Println("\n making request to:", req.Endpoint)
	info, _ := configureRedirectToWSS(c)

	newReq, err := http.NewRequestWithContext(req.Context, req.Method, req.Endpoint, req.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error in creating new request:\n: %w", err)
	}

	if c == nil {
		panic("api.LoginRoute:client pointer is nil")
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

	fmt.Printf("\nRedirect-link: %s\n", info.url.String())
	return serverRes, info, nil
}
