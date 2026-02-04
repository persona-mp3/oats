package common

import "net/url"

var WelcomeEndpoint = "/welcome"
var LoginEndpoint = "/login"
var AuthServer = "http://localhost:8000"

var ChatEvent = 1
var FindFlag = 2

type Credentials struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type RedirectInfo struct {
	Location *url.URL
}
