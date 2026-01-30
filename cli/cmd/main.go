package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	POST_METHOD   = "POST"
	addr          = "http://localhost:8990"
	docsRoute     = "/docs"
	helpRoute     = "/help-me"
	loginRoute    = "/login"
	registerRoute = "/register"
	welcomeRoute  = "/welcome"
)

func newClient() *http.Client {
	return &http.Client{}
}

type Req struct {
	Method   string
	Endpoint string
	Context  context.Context
	Body     io.Reader
}

type Res struct {
	Code int
	Body []byte
}

func (r *Req) makeRequest(c *http.Client) (*Res, error) {
	newReq, err := http.NewRequestWithContext(r.Context, r.Method, r.Endpoint, r.Body)
	if err != nil {
		return nil, fmt.Errorf("error in creating request:%w", err)
	}

	if c == nil {
		log.Fatal("client is nill")
	}

	newReq.Header.Set("Content-Type", "application/json")
	res, err := c.Do(newReq)
	if err != nil {
		return nil, fmt.Errorf("error in contacting server:%w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error in reading response body:%w", err)
	}

	return &Res{res.StatusCode, body}, nil
}

func parseCmdLine() string {
	// later on, add validation to check if the endpoint actually exists
	var ep string
	flag.StringVar(&ep, "ep", "welcome", "endpoint to visit, default is welcome")
	flag.Parse()

	if len(ep) < 5 {
		fmt.Println("contacting default route")
		ep = welcomeRoute
	}

	return ep
}

type UserCredentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func parseLoginCredentials() *UserCredentials {
	var userName string
	var passwd string
	var email string

	for {
		fmt.Print("\n provide userName: ")
		fmt.Scanf("%s", &userName)

		fmt.Print("\n password: ")
		fmt.Scanf("%s", &passwd)

		fmt.Print("\n email: ")
		fmt.Scanf("%s", &email)

		if len(userName) < 3 || len(passwd) < 3 || len(email) < 3 {
			fmt.Println("puhlease provide good creds")
			fmt.Println(userName, passwd, email)
			continue
		}

		break
	}

	return &UserCredentials{userName, passwd, email}
}

func toJson(data any) ([]byte, error) {
	content, err := json.Marshal(data)
	if err != nil {
		return []byte{}, fmt.Errorf("error in converting to_json: %w", err)
	}
	return content, nil
}

func main() {
	tgtEndpoint := parseCmdLine()

	newReq := &Req{
		Method:   "GET",
		Endpoint: addr + tgtEndpoint,
		Body:     nil,
	}

	if tgtEndpoint == registerRoute || tgtEndpoint == loginRoute {
		fmt.Println("hitting /login, provide credentials")

		creds := parseLoginCredentials()
		data, err := toJson(creds)

		if err != nil {
			log.Fatal(err)
		}

		newReq.Method = POST_METHOD
		newReq.Body = bytes.NewReader(data)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newReq.Context = ctx

	c := newClient()
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		res := Res{}
		res.Code = req.Response.StatusCode
		url, err := req.Response.Location()
		if err != nil {
			return fmt.Errorf("occured in getting Location from redirect\n:%w", err)
		}

		fmt.Println("irl->", url)
		fmt.Printf("redirect-response -> %+v+\n", req.Response)
		return http.ErrUseLastResponse
	}

	res, err := newReq.makeRequest(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("response from server:")
	fmt.Printf("\n%s\n", res.Body)
}
