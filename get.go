package scout

import (
	"context"
	"net/http"
)

func (c *Client) PleaseGet(ctx context.Context, path string) (*Response, error) {
	return c.Do(ctx, http.MethodGet, path, nil)
}
