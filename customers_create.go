package resellerclub

import "strconv"

type CustomerCreateParams struct {
	// Required. Username for the Customer Account. Username should be an email address.
	Username string

	// Required. Password for the Customer Account.
	// Allowed Password length is 9 to 16 characters. It should contain at least:
	//  1 lowercase character
	//  1 uppercase character
	//  1 number
	//  1 special character. (Allowed special characters are: ~*!@$#%_+.?:,{})
	Password string

	// Required. Name of the Customer
	Name string

	// Required. Name of the Customer's company
	Company string

	// Required. Address line 1 of the Customer's address
	AddressLine1 string

	// Required. City
	City string

	// Required. State. In case the State information is not available, you need to pass the value for this parameter as Not Applicable.
	State string

	// Optional. This parameter needs to be included if the State information is not available.
	// Mention an appropriate value for this parameter.
	OtherState string

	// Required. Country Code as per ISO 3166-1 alpha-2
	Country string

	// Required. ZIP code
	Zipcode string

	// Required. Telephone number Country Code
	PhoneCC string

	// Required. Phone number
	Phone string

	// Required. Language Code as per ISO
	LangPref string

	// Optional. Address line 2 of the Customer's address
	AddressLine2 string

	// Optional. Address line 3 of the Customer's address
	AddressLine3 string

	// Optional. Alternate phone country code
	AltPhoneCC string

	// Optional. Alternate phone number
	AltPhone string

	// Optional. Fax number country code
	FaxCC string

	// Optional. Fax number
	Fax string

	// Optional. Mobile country code
	MobileCC string

	// Optional. Mobile number
	Mobile string

	// Optional. In case of a US based customer, consent is required to receive renewal reminder SMSes
	SMSConsent bool

	// Optional. VAT ID for EU VAT
	VATID string

	// Optional. Accept Terms and Conditions and Privacy Policy to create an account
	AcceptPolicy bool

	// Optional. In case of EEA (European Economic Area) countries capture consent to receive marketing emails
	MarketingEmailConsent bool
}

// Create creates a new customer.
// https://manage.resellerclub.com/kb/answer/804
func (customers *Customers) Create(params *CustomerCreateParams) (int64, error) {
	if params == nil {
		return 0, ErrMissingParams
	}

	u := customers.url("/v2/signup.json")
	q := u.Query()

	q.Set("username", params.Username)
	q.Set("passwd", params.Password)
	q.Set("name", params.Name)
	q.Set("company", params.Company)
	q.Set("address-line-1", params.AddressLine1)
	q.Set("city", params.City)
	q.Set("state", params.State)
	q.Set("country", params.Country)
	q.Set("zipcode", params.Zipcode)
	q.Set("phone-cc", params.PhoneCC)
	q.Set("phone", params.Phone)
	q.Set("lang-pref", params.LangPref)

	if len(params.OtherState) > 0 {
		q.Set("other-state", params.OtherState)
	}
	if len(params.AddressLine2) > 0 {
		q.Set("address-line-2", params.AddressLine2)
	}
	if len(params.AddressLine3) > 0 {
		q.Set("address-line-3", params.AddressLine3)
	}
	if len(params.AltPhoneCC) > 0 {
		q.Set("alt-phone-cc", params.AltPhoneCC)
	}
	if len(params.AltPhone) > 0 {
		q.Set("alt-phone", params.AltPhone)
	}
	if len(params.FaxCC) > 0 {
		q.Set("fax-cc", params.FaxCC)
	}
	if len(params.Fax) > 0 {
		q.Set("fax", params.Fax)
	}
	if len(params.MobileCC) > 0 {
		q.Set("mobile-cc", params.MobileCC)
	}
	if len(params.Mobile) > 0 {
		q.Set("mobile", params.Mobile)
	}
	if params.SMSConsent {
		q.Set("sms-consent", strconv.FormatBool(params.SMSConsent))
	}
	if len(params.VATID) > 0 {
		q.Set("vat-id", params.VATID)
	}
	if params.AcceptPolicy {
		q.Set("accept-policy", strconv.FormatBool(params.AcceptPolicy))
	}
	if params.MarketingEmailConsent {
		q.Set("marketing-email-consent", strconv.FormatBool(params.MarketingEmailConsent))
	}

	u.RawQuery = q.Encode()

	return customers.client.postInt64(u.String())
}
