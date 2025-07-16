package scout

type Client struct {
	Url           string
	Authorization string
	// more options coming soon...
}

type Response struct {
	StatusCode int    // 200
	Status     string // 200 OK
}

// TO-DO: Build Response Type
func NewClient(url string) *Client {
	return &Client{Url: url}
}
