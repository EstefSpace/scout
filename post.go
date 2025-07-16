package scout

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) PleasePostWithJSON(ctx context.Context, path string, data interface{}) (*Response, error) {
	body, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewReader(body)

	resp, err := c.Do(ctx, http.MethodPost, path, reqBody)
	if err == nil {
		resp.Header.Add("Content-Type", "Application/Json")
	}

	return resp, err
}
