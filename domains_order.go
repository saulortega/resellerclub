package resellerclub

type OrderContact struct {
	Company       string   `json:"company"`
	Address1      string   `json:"address1"`
	TelNo         string   `json:"telno"`
	TelNoCC       string   `json:"telnocc"`
	ContactID     Int64    `json:"contactid"`
	Type          string   `json:"type"`
	ContactType   []string `json:"contacttype"`
	CustomerID    Int64    `json:"customerid"`
	Country       string   `json:"country"`
	ParentKey     string   `json:"parentkey"`
	ContactStatus string   `json:"contactstatus"`
	State         string   `json:"state"`
	EmailAddr     string   `json:"emailaddr"`
	City          string   `json:"city"`
	Name          string   `json:"name"`
	Zip           string   `json:"zip"`
}

// GetOrderID gets the Order Id of a Registered domain name.
// https://manage.resellerclub.com/kb/answer/763
func (domains *Domains) GetOrderID(domainName string) (int64, error) {
	u := domains.url("/orderid.json")
	q := u.Query()

	q.Set("domain-name", domainName)

	u.RawQuery = q.Encode()

	id, err := domains.client.getInt64(u.String())
	if err != nil {
		return 0, err
	}

	return id, nil
}
