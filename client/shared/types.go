package shared

import (
	"net/url"
)

var (
	WelcomeEndpoint = "/welcome"
)

type Credentials struct {
	Username string
	Password string
}

type RedirectInfo struct {
	Url        *url.URL
	StatusCode int
}
