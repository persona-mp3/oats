package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func (req *Req) GenericRequest(c *http.Client) (*GenericRes, error) {
	newReq, err := http.NewRequestWithContext(req.Context, req.Method, req.Endpoint, req.Body)
	if err != nil {
		return nil, fmt.Errorf("error in creating request:%w", err)
	}

	if c == nil {
		log.Fatalf("[panic] client pointer is nill")
	}

	newReq.Header.Set("Content-Type", "application/json")
	res, err := c.Do(newReq)
	if err != nil {
		return nil, fmt.Errorf("could not contact server:%w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read content:%w", err)
	}

	return &GenericRes{res.StatusCode, body, nil}, nil

}
