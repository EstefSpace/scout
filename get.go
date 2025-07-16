package scout

import (
	"net/http"
)

func (c *Client) PleaseGetThis(route string) (*Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url+route, nil)

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
