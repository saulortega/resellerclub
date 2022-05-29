package resellerclub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

const (
	DomainStatusAvailable               = "available"
	DomainStatusRegisteredThroughUs     = "regthroughus"
	DomainStatusRegisteredThroughOthers = "regthroughothers"
	DomainStatusUnknown                 = "unknown"
)

type DomainAvailabilityResponse struct {
	Domain   string `json:"domain"`
	Classkey string `json:"classkey"`
	Status   string `json:"status"`
}

// CheckAvailability checks domains' availability.
// https://manage.resellerclub.com/kb/answer/764
func (domains *Domains) CheckAvailability(domainNames []string, tlds []string) ([]*DomainAvailabilityResponse, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, somethingWentWrong(string(body))
	}

	mapResp := map[string]interface{}{}
	err = json.Unmarshal(body, &mapResp)
	if err != nil {
		return nil, somethingWentWrong(string(body))
	}

	err = checkResponseError(mapResp)
	if err != nil {
		return nil, err
	}

	domainsAvailability := []*DomainAvailabilityResponse{}
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

		domainsAvailability = append(domainsAvailability, &DomainAvailabilityResponse{
			Domain:   domain,
			Classkey: classkey,
			Status:   status,
		})
	}

	if len(domainsAvailability) == 0 {
		return nil, somethingWentWrong(mapResp)
	}

	sort.Slice(domainsAvailability, func(i, j int) bool {
		return domainsAvailability[i].Domain < domainsAvailability[j].Domain
	})

	return domainsAvailability, nil
}
