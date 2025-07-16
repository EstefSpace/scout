package scout

import (
	"context"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	Url        string
	headers    map[string]string
	// more options coming soon...
}

type Response struct {
	*http.Response
}

func NewClient(url string) *Client {
	return &Client{Url: url}
}

func NewResponse(r *http.Response) *Response {
	return &Response{Response: r}
}

func (c *Client) PleaseAddHeader(key, value string) *Client {
	c.headers[key] = value
	return c
}

func (c *Client) Do(ctx context.Context, method, path string, body io.Reader) (*Response, error) {
	url := c.Url + path

	req, err := http.NewRequestWithContext(ctx, method, url, body)

	if err != nil {
		return nil, err
	}

	// set alls headers

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(httpResp), err
}
