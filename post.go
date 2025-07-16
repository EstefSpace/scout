package scout

import (
	"io"
	"net/http"
)

func (c *Client) PleasePostThis(route string, info io.Reader) (*Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", c.Url+route, info)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.Authorization)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: resp.StatusCode, Status: resp.Status}, nil
}
