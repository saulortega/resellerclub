package resellerclub

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Domains struct {
	client *Client
}

type DomainAvailability struct {
	Domain   string `json:"domain"`
	Classkey string `json:"classkey"`
	Status   string `json:"status"`
}

func (domains *Domains) url(path string) *url.URL {
	u := domains.client.url("/domains")
	u.Path += path
	return u
}

func (domains *Domains) CheckAvailability(domainNames []string, tlds []string) ([]DomainAvailability, error) {
	u := domains.url("/available.json")
	q := u.Query()

	q["domain-name"] = domainNames
	q["tlds"] = tlds
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	mapResp := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&mapResp)
	if err != nil {
		return nil, err
	}

	err = checkResponseError(mapResp)
	if err != nil {
		return nil, err
	}

	domainsAvailability := []DomainAvailability{}
	for domain, data := range mapResp {
		d, ok := data.(map[string]interface{})
		if !ok {
			return nil, somethingWentWrong(mapResp)
		}

		status, ok := d["status"].(string)
		if !ok {
			return nil, somethingWentWrong(mapResp)
		}

		classkey, _ := d["classkey"].(string)

		domainsAvailability = append(domainsAvailability, DomainAvailability{
			Domain:   domain,
			Classkey: classkey,
			Status:   status,
		})
	}

	if len(domainsAvailability) == 0 {
		return nil, somethingWentWrong(mapResp)
	}

	return domainsAvailability, nil
}
