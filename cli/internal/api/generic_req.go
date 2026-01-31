package api

import (
	"fmt"
	"io"
	"net/http"
)

func (req *Req) GenericRequest(c *http.Client) (*GenericRes, error) {
	newReq, err := http.NewRequestWithContext(req.Context, req.Method, req.Endpoint, req.Body)
	if err != nil {
		return nil, fmt.Errorf("error in creating request:%w", err)
	}

	if c == nil {
		panic("client is nill")
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

	return &GenericRes{res.StatusCode, body, nil}, nil

}
