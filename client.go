package resellerclub

import "net/url"

type Client struct {
	userID string
	key    string

	Domains *Domains
}

func New(userID string, key string) *Client {
	client := &Client{
		userID: userID,
		key:    key,
	}

	client.Domains = &Domains{client}

	return client
}

func (client *Client) url(path string) *url.URL {
	u, _ := url.Parse(endpoint)
	q := u.Query()
	q.Set("auth-userid", client.userID)
	q.Set("api-key", client.key)
	u.Path += path
	u.RawQuery = q.Encode()

	return u
}
