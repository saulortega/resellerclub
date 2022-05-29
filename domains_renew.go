package resellerclub

import (
	"strconv"
	"time"
)

type DomainRenewParams struct {
	// Required. Order Id of the Domain Registration Order that you want to Renew.
	OrderID int64

	// Required. Number of years for which you want to Renew this Order.
	Years int

	// Required. Current Expiry Date of the Order in epoch time format.
	ExpDate time.Time

	// Optional. Renews the Privacy Protection service for the domain name.
	// Privacy Protection is not supported for the following TLDs:
	//  .ASIA, .AU, .CA, .CN, .ORG.CO, .MIL.CO, .GOV.CO, .EDU.CO, .DE, .ES, .EU, .IN, .NL, .NZ, .PRO, .RU, .SX, .TEL, .UK, .US.
	PurchasePrivacy bool

	// Required. Enables / Disables the Auto Renewal setting for the domain name.
	AutoRenew bool

	// Required. This will decide how the Customer Invoice will be handled. Set any of below mentioned Invoice Options for your Customer:
	//  NoInvoice: This will not raise any Invoice. The Order will be renewed.
	//  PayInvoice: This will raise an Invoice and:
	//   - if there are sufficient funds in the Customer's Debit Account, then the Invoice will be paid and the Order will be renewed.
	//   - if there are insufficient funds in the Customer's Debit Account, then the Order will remain pending in the system.
	//  KeepInvoice: This will raise an Invoice for the Customer to pay later. The Order will be renewed.
	//  OnlyAdd: This will raise an Invoice for the Customer to pay later. The renewal action request will remain pending.
	InvoiceOption string

	// Optional. Discount amount for the order value.
	DiscountAmount float64

	// Optional. Purchase Premium DNS Service.
	PurchasePremiumDNS bool
}

type DomainRenewResponse struct {
	DomainCommonResponse

	// Privacy Protection Details
	PrivacyDetails DomainCommonResponse `json:"privacydetails"`
}

type resDomainRenewResponse struct {
	errorResponse
	DomainRenewResponse
}

// Renew renews a domain name.
// https://manage.resellerclub.com/kb/answer/746
func (domains *Domains) Renew(params *DomainRenewParams) (*DomainRenewResponse, error) {
	if params == nil {
		return nil, ErrMissingParams
	}

	u := domains.url("/renew.json")
	q := u.Query()

	q.Set("order-id", strconv.FormatInt(params.OrderID, 10))
	q.Set("years", strconv.Itoa(params.Years))
	q.Set("exp-date", strconv.FormatInt(params.ExpDate.Unix(), 10))
	q.Set("auto-renew", strconv.FormatBool(params.AutoRenew))
	q.Set("invoice-option", params.InvoiceOption)

	if params.PurchasePrivacy {
		q.Set("purchase-privacy", strconv.FormatBool(params.PurchasePrivacy))
	}
	if params.DiscountAmount > 0 {
		q.Set("discount-amount", strconv.FormatFloat(params.DiscountAmount, 'f', -1, 64))
	}
	if params.PurchasePremiumDNS {
		q.Set("purchase-premium-dns", strconv.FormatBool(params.PurchasePremiumDNS))
	}

	u.RawQuery = q.Encode()

	var res = resDomainRenewResponse{}
	err := domains.client.post(u.String(), &res)
	if err != nil {
		return nil, err
	}

	return &res.DomainRenewResponse, nil
}
