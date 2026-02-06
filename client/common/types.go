package common

import "net/url"

var WelcomeEndpoint = "/welcome"
var LoginEndpoint = "/login"
var AuthServer = "http://localhost:8000"

var SingleChat = "i"
var MulChat = "a"
var QuitChat = "q"

var ChatEvent = 1
var FindFlag = 2

type Credentials struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type RedirectInfo struct {
	Location *url.URL
}

type Event struct {
	Name  int
	Value string
}

type Friend struct {
	Name     string `json:"name"`
	LastSeen string `json:"lastSeen"`
}

type Message struct {
	Dest    string `json:"dest"`
	From    string `json:"from"`
	Time    string `json:"time"`
	Message string `json:"message"`
}
