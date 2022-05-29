package resellerclub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type DomainSearchParams struct {
	// Required. Number of Orders to be fetched. This should be a value between 10 to 500.
	NoOfRecords int

	// Required. Page number for which details are to be fetched.
	PageNo int

	// Optional. One or more parameters by which you want to sort the Orders.
	// Values can be orderid, customerid, endtime, timestamp, entitytypeid, creationtime or creationdt.
	// You can sort your results in the ascending or the descending order by passing the asc or desc parameter in the call.
	// By default, the orders are sorted in the ascending order. Example: &order-by=timestamp desc
	// Example:
	//  If page-no is 1, no-of-records is 30 and order-by is orderid;
	//  it will fetch the first 30 Orders which suit the search criteria sorted by orderid.
	//  If page-no is 2, no-of-records is 30 and order-by is order id;
	//  it will fetch the second batch of 30 Orders sorted by orderid.
	OrderBy []string

	// Optional. Order Id(s) of the Domain Registration Order(s) whose details need to be fetched.
	OrderIDs []int64

	// Optional. Reseller Id(s) whose Orders need to be fetched.
	ResellerIDs []int64

	// Optional. Customer Id(s) whose Orders need to be fetched.
	CustomerIDs []int64

	// Optional. Whether Sub-Reseller Orders need to be fetched or not.
	ShowChildOrders bool

	// Optional. Product keys of the TLDs.
	ProductKey []string

	// Optional. Status of the Order, namely, InActive, Active, Suspended, Pending Delete Restorable, Deleted, Archived, Pending Verification or Failed Verification
	//  Deleted - to be used for searching Orders deleted in the past 30 days
	//  Archived - to be used for searching Orders deleted more than 30 days ago
	//  Pending Verification - to be used for searching Orders for which the Registrant Contact email address verification is pending
	//  Failed Verification - to be used for searching Orders which have been deactivated due to non-verification of the Registrant Contact email address
	Status []string

	// Optional. Name of the Domain.
	DomainName string

	// Optional. Privacy Protection status of the Order, namely, true, false or na
	//  true - to be used for searching Orders for which Privacy Protection is enabled
	//  false - to be used for searching Orders for which Privacy Protection is disabled
	//  na - to be used for searching Orders for which Privacy Protection is not applicable (not allowed). The search results will include Inactive Orders as well.
	PrivacyEnabled string

	// Optional. UNIX TimeStamp for listing of Domain Registration Orders whose Creation Date is greater than creation-date-start.
	CreationDateStart time.Time

	// Optional. UNIX TimeStamp for listing of Domain Registration Orders whose Creation Date is less than creation-date-end.
	CreationDateEnd time.Time

	// Optional. UNIX TimeStamp for listing of Domain Registration Orders whose Expiry Date is greater than expiry-date-start.
	ExpiryDateStart time.Time

	// Optional. UNIX TimeStamp for listing of Domain Registration Orders whose Expiry Date is less than expiry-date-end.
	ExpiryDateEnd time.Time
}

type DomainSearchResponseItem struct {
	OrderID           Int64  `json:"orders.orderid"`            // Order ID
	CustomerID        Int64  `json:"entity.customerid"`         // Customer ID
	EntityID          Int64  `json:"entity.entityid"`           // Entity ID
	Autorenew         Bool   `json:"orders.autorenew"`          // Autorenew Status
	EndTime           Time   `json:"orders.endtime"`            // Expiry (at the Registry) Timestamp
	ResellerLock      Bool   `json:"orders.resellerlock"`       // Reseller Lock Status
	Timestamp         Time   `json:"orders.timestamp"`          // Timestamp for the Last Modification Action on the Order
	CustomerLock      Bool   `json:"orders.customerlock"`       // Customer Lock Status
	EntityTypeID      Int64  `json:"entity.entitytypeid"`       // Entity type ID
	CurrentStatus     string `json:"entity.currentstatus"`      // Current Order Status: InActive, Active, Suspended, Pending Delete Restorable, Deleted, Archived, Pending Verification or Failed Verification
	EntityTypeKey     string `json:"entitytype.entitytypekey"`  // Product Key
	TransferLock      Bool   `json:"orders.transferlock"`       // Transfer Lock Status
	CreationTime      Time   `json:"orders.creationtime"`       // Order Creation (at the Registry) Timestamp
	PrivacyProtection Bool   `json:"orders.privacyprotection"`  // Privacy Protection Status
	EntityTypeName    string `json:"entitytype.entitytypename"` // Product Name
	Creationdt        Time   `json:"orders.creationdt"`         // Order Addition Timestamp
	Description       string `json:"entity.description"`        // Domain Name
}

// Search gets a list of Domain Registration Orders matching the search criteria, along with the details.
// https://manage.resellerclub.com/kb/answer/771
func (domains *Domains) Search(params *DomainSearchParams) ([]*DomainSearchResponseItem, error) {
	if params == nil {
		return nil, ErrMissingParams
	}

	u := domains.url("/search.json")
	q := u.Query()

	q.Set("no-of-records", strconv.Itoa(params.NoOfRecords))
	q.Set("page-no", strconv.Itoa(params.PageNo))
	if len(params.OrderBy) > 0 {
		q["order-by"] = params.OrderBy
	}
	if len(params.OrderIDs) > 0 {
		for _, i := range params.OrderIDs {
			q.Add("order-id", strconv.FormatInt(i, 10))
		}
	}
	if len(params.ResellerIDs) > 0 {
		for _, i := range params.ResellerIDs {
			q.Add("reseller-id", strconv.FormatInt(i, 10))
		}
	}
	if len(params.CustomerIDs) > 0 {
		for _, i := range params.CustomerIDs {
			q.Add("customer-id", strconv.FormatInt(i, 10))
		}
	}
	if params.ShowChildOrders {
		q.Set("show-child-orders", "true")
	}
	if len(params.ProductKey) > 0 {
		q["product-key"] = params.ProductKey
	}
	if len(params.Status) > 0 {
		q["status"] = params.Status
	}
	if len(params.DomainName) > 0 {
		q.Set("domain-name", params.DomainName)
	}
	if len(params.PrivacyEnabled) > 0 {
		q.Set("privacy-enabled", params.PrivacyEnabled)
	}
	if !params.CreationDateStart.IsZero() {
		q.Set("creation-date-start", strconv.FormatInt(params.CreationDateStart.Unix(), 10))
	}
	if !params.CreationDateEnd.IsZero() {
		q.Set("creation-date-end", strconv.FormatInt(params.CreationDateEnd.Unix(), 10))
	}
	if !params.ExpiryDateStart.IsZero() {
		q.Set("expiry-date-start", strconv.FormatInt(params.ExpiryDateStart.Unix(), 10))
	}
	if !params.ExpiryDateEnd.IsZero() {
		q.Set("expiry-date-end", strconv.FormatInt(params.ExpiryDateEnd.Unix(), 10))
	}

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

	var errRes errorResponse
	err = json.Unmarshal(body, &errRes)
	if err != nil {
		return nil, somethingWentWrong(string(body))
	}

	if err = errRes.Err(); err != nil {
		return nil, err
	}

	var res map[string]json.RawMessage
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, somethingWentWrong(string(body))
	}

	var items []*DomainSearchResponseItem
	for k, v := range res {
		if k[0] < '1' || k[0] > '9' {
			continue
		}

		var item DomainSearchResponseItem
		err = json.Unmarshal(v, &item)
		if err != nil {
			return nil, somethingWentWrong(v)
		}

		items = append(items, &item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Description < items[j].Description
	})

	return items, nil
}
