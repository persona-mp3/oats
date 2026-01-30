package api

import (
	"context"
	"io"
	"net/url"
)

type NewReq struct {
	Artist string
}

type Req struct {
	Method   string
	Endpoint string
	Context  context.Context
	Body     io.Reader
}

type GenericRes struct {
	StatusCode int
	Content    []byte
	Data       any
}

type UserCredentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RedirectInfo struct {
	url        *url.URL
	statusCode int
}
