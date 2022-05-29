package resellerclub

import (
	"fmt"
	"strconv"
)

// https://manage.resellerclub.com/kb/answer/752
type DomainRegisterParams struct {
	// Required. Domain name that you need to Register.
	DomainName string

	// Required. Number of years for which you wish to Register this domain name.
	Years int

	// Required. The Name Servers of the domain name.
	NS []string

	// Required. The Customer for whom you wish to Register this domain name.
	CustomerID int64

	// Required. The Registrant Contact of the domain name.
	RegContactID int64

	// Required. The Administrative Contact of the domain name.
	// Pass -1 for the TLDs .EU, .NZ, .RU, .UK.
	AdminContactID int64

	// Required. The Technical Contact of the domain name.
	// Pass -1 for the TLDs .EU, .FR, .NZ, .RU, .UK.
	TechContactID int64

	// Required. The Billing Contact of the domain name.
	// Pass -1 for the TLDs .BERLIN, .CA, .EU, .FR, .NL, .NZ, .RU, .UK, .LONDON.
	BillingContactID int64

	// Required. This will decide how the Customer Invoice will be handled. Set any of below mentioned Invoice Options for your Customer:
	//  NoInvoice: This will not raise any Invoice. The Order will be executed.
	//  PayInvoice: This will raise an Invoice and:
	//   - if there are sufficient funds in the Customer's Debit Account, then the Invoice will be paid and the Order will be executed.
	//   - if there are insufficient funds in the Customer's Debit Account, then the Order will remain pending in the system.
	//  KeepInvoice: This will raise an Invoice for the Customer to pay later. The Order will be executed.
	//  OnlyAdd: This will raise an Invoice for the Customer to pay later. The registration action request will remain pending.
	InvoiceOption string

	// Optional. Adds the Privacy Protection service for the domain name.
	// Privacy Protection is not supported for the following TLDs (extensions):
	//  .ASIA, .AU, .CA, .CL, .CN, .ORG.CO, .MIL.CO, .GOV.CO, .EDU.CO, .DE, .ES, .EU, .FR, .IN, .NL, .NZ, .PRO, .RU, .SX, .TEL, .UK, .US.
	PurchasePrivacy bool

	// Optional. Enables / Disables the Privacy Protection setting for the domain name.
	ProtectPrivacy bool

	// Required. Enables / Disables the Auto Renewal setting for the domain name.
	AutoRenew bool

	// Optional. Mapping key of the extra details needed to register a domain name.
	// See more details of attr-name/attr-value in https://manage.resellerclub.com/kb/answer/752
	ExtraAttrs map[string]string

	// Optional. Discount amount for the order value.
	DiscountAmount float64

	// Optional. Purchase Premium DNS service.
	PurchasePremiumDNS bool
}

// https://manage.resellerclub.com/kb/answer/752
type DomainRegisterResponse struct {
	DomainCommonResponse

	// Privacy Protection Details
	PrivacyDetails DomainCommonResponse `json:"privacydetails"`
}

type resDomainRegisterResponse struct {
	errorResponse
	DomainRegisterResponse
}

// Register registers a domain name.
// https://manage.resellerclub.com/kb/answer/752
func (domains *Domains) Register(params *DomainRegisterParams) (*DomainRegisterResponse, error) {
	if params == nil {
		return nil, ErrMissingParams
	}

	u := domains.url("/register.json")
	q := u.Query()

	q.Set("domain-name", params.DomainName)
	q.Set("years", strconv.Itoa(params.Years))
	q["ns"] = params.NS
	q.Set("customer-id", strconv.FormatInt(params.CustomerID, 10))
	q.Set("reg-contact-id", strconv.FormatInt(params.RegContactID, 10))
	q.Set("admin-contact-id", strconv.FormatInt(params.AdminContactID, 10))
	q.Set("tech-contact-id", strconv.FormatInt(params.TechContactID, 10))
	q.Set("billing-contact-id", strconv.FormatInt(params.BillingContactID, 10))
	q.Set("invoice-option", params.InvoiceOption)
	q.Set("auto-renew", strconv.FormatBool(params.AutoRenew))

	if params.PurchasePrivacy {
		q.Set("purchase-privacy", strconv.FormatBool(params.PurchasePrivacy))
	}
	if params.ProtectPrivacy {
		q.Set("protect-privacy", strconv.FormatBool(params.ProtectPrivacy))
	}
	if params.DiscountAmount > 0 {
		q.Set("discount-amount", strconv.FormatFloat(params.DiscountAmount, 'f', -1, 64))
	}
	if params.PurchasePremiumDNS {
		q.Set("purchase-premium-dns", strconv.FormatBool(params.PurchasePremiumDNS))
	}

	var i = 1
	for k, v := range params.ExtraAttrs {
		q.Set(fmt.Sprintf("attr-name%v", i), k)
		q.Set(fmt.Sprintf("attr-value%v", i), v)
		i++
	}

	u.RawQuery = q.Encode()

	var res = resDomainRegisterResponse{}
	err := domains.client.post(u.String(), &res)
	if err != nil {
		return nil, err
	}

	return &res.DomainRegisterResponse, nil
}
