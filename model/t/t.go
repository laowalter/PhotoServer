package main

func main() {
}

type Client struct {
	authKey string // private authKey set in constructors
}

type ClientOption func(c *Client)

func (c *Client) AuthKey() string {
	if c.authKey != "" {
		return c.authKey
	}
	return "default-auth-key"
}

func NewClientWithAuth(key string) *Client { // 'With' Constructor
	return &Client{key}
}

func NewClient(opts ...ClientOption) *Client {
	client := &Client{}
	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithAuth(key string) ClientOption {
	// this is the ClientOption function type
	return func(c *Client) {
		c.authKey = key
	}
}
