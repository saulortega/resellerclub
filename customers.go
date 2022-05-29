package resellerclub

import (
	"net/url"
)

type Customers struct {
	client *Client
}

func (customers *Customers) url(path string) *url.URL {
	u := customers.client.url("/customers")
	u.Path += path
	return u
}
