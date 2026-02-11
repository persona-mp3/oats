package shared

import (
	"net/url"
)

type Credentials struct {
	Username string
	Password string
}

type RedirectInfo struct {
	Url        *url.URL
	StatusCode int
}
