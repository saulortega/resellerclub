package resellerclub

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	userID string
	key    string

	Domains   *Domains
	Customers *Customers
}

func New(userID string, key string) *Client {
	client := &Client{
		userID: userID,
		key:    key,
	}

	client.Domains = &Domains{client}
	client.Customers = &Customers{client}

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

func (client *Client) get(url string, target errorChecker) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return client.request(resp, target)
}

func (client *Client) post(url string, target errorChecker) error {
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return client.request(resp, target)
}

func (client *Client) request(resp *http.Response, target errorChecker) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return somethingWentWrong(string(body))
	}

	if len(body) > 0 && body[0] != '{' {
		return errors.New(string(body))
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return somethingWentWrong(string(body))
	}

	if err = target.Err(); err != nil {
		return err
	}

	return nil
}

func (client *Client) getInt64(url string) (int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return client.requestInt64(resp)
}

func (client *Client) postInt64(url string) (int64, error) {
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return client.requestInt64(resp)
}

func (client *Client) requestInt64(resp *http.Response) (int64, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, somethingWentWrong(string(body))
	}

	var i int64
	err = json.Unmarshal(body, &i)
	if err != nil {
		var e errorResponse
		err = json.Unmarshal(body, &e)
		if err != nil {
			return 0, somethingWentWrong(string(body))
		}

		if err = e.Err(); err != nil {
			return 0, err
		}

		return 0, somethingWentWrong(string(body))
	}

	return i, nil
}
