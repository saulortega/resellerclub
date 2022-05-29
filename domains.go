package resellerclub

import (
	"net/url"
)

type Domains struct {
	client *Client
}

type DomainCommonResponse struct {
	Description             string  `json:"description"`             // Domain Name
	EntityID                Int64   `json:"entityid"`                // Order ID of the Domain Order
	ActionType              string  `json:"actiontype"`              // Action Type
	ActionTypeDesc          string  `json:"actiontypedesc"`          // Description of the Privacy Protection Action
	EaqID                   Int64   `json:"eaqid"`                   // Action ID of the Privacy Protection Action
	ActionStatus            string  `json:"actionstatus"`            // Action Status
	ActionStatusDesc        string  `json:"actionstatusdesc"`        // Description of the Action Status
	InvoiceID               Int64   `json:"invoiceid"`               // Invoice ID of the Invoice
	SellingCurrencySymbol   string  `json:"sellingcurrencysymbol"`   // Selling Currency of the Reseller
	SellingAmount           Float64 `json:"sellingamount"`           // Transaction Amount in the Selling Currency
	UnutilisedSellingAmount Float64 `json:"unutilisedsellingamount"` // Unutilized Transaction Amount in the Selling Currency
	DiscountAmount          Float64 `json:"discount-amount"`         // Discount Amount
	CustomerID              Int64   `json:"customerid"`              // Customer ID associated with the Domain Order
}

func (domains *Domains) url(path string) *url.URL {
	u := domains.client.url("/domains")
	u.Path += path
	return u
}
