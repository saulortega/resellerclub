package resellerclub

import "strconv"

type OrderDetailsOption string

const (
	OrderDetailsOptionAll                      OrderDetailsOption = "All"
	OrderDetailsOptionOrderDetails             OrderDetailsOption = "OrderDetails"
	OrderDetailsOptionContactIds               OrderDetailsOption = "ContactIds"
	OrderDetailsOptionRegistrantContactDetails OrderDetailsOption = "RegistrantContactDetails"
	OrderDetailsOptionAdminContactDetails      OrderDetailsOption = "AdminContactDetails"
	OrderDetailsOptionTechContactDetails       OrderDetailsOption = "TechContactDetails"
	OrderDetailsOptionBillingContactDetails    OrderDetailsOption = "BillingContactDetails"
	OrderDetailsOptionNsDetails                OrderDetailsOption = "NsDetails"
	OrderDetailsOptionDomainStatus             OrderDetailsOption = "DomainStatus"
	OrderDetailsOptionDNSSECDetails            OrderDetailsOption = "DNSSECDetails"
	OrderDetailsOptionStatusDetails            OrderDetailsOption = "StatusDetails"
)

type DomainGetOrderDetailsResponse struct {
	// Order ID
	OrderID Int64 `json:"orderid"`

	// Order Description
	Description string `json:"description"`

	// Domain Name
	DomainName string `json:"domainname"`

	// Current Order Status under the System.
	// Value will be InActive, Active, Suspended, Pending Delete Restorable, Deleted or Archived
	CurrentStatus string `json:"currentstatus"`

	// Lock/Hold on the domain name at the Registry.
	// Value(s) will be resellersuspend, resellerlock and/or transferlock
	OrderStatus []string `json:"orderstatus"`

	// Lock/Hold on the domain name in the System.
	// Value(s) will be sixtydaylock and/or renewhold
	DomainStatus []string `json:"domainstatus"`

	// Product Category
	ProductCategory string `json:"productcategory"`

	// Product Key
	ProductKey string `json:"productkey"`

	// Order Creation (at the Registry) Date
	CreationTime Time `json:"creationtime"`

	// Registrant Contact Email Address Verification Status.
	// Value will be Verified, Pending or Suspended
	RaaVerificationStatus string `json:"raaVerificationStatus"`

	// Expiry Date (at the Registry)
	EndTime Time `json:"endtime"`

	// Whether Order belongs to a Customer directly under the Reseller
	IsImmediateReseller Bool `json:"isImmediateReseller"`

	// Reseller Chain by RID
	ParentKey string `json:"parentkey"`

	// Customer ID Associated with the Order
	CustomerID Int64 `json:"customerid"`

	// Number of Name Servers associated with the Domain Name
	NoOfNameServers Int64 `json:"noOfNameServers"`

	// Name Server 1
	NS1 string `json:"ns1"`

	// Name Server 2
	NS2 string `json:"ns2"`

	// Domain Secret
	DomainSecret string `json:"domsecret"`

	// Whether Order Suspended due to Expiry
	IsOrderSuspendedUponExpiry Bool `json:"isOrderSuspendedUponExpiry"`

	// Whether Order Suspended by Parent Reseller
	OrderSuspendedByParent Bool `json:"orderSuspendedByParent"`

	// Whether Privacy Protection allowed for the Product Type
	PrivacyProtectedAllowed Bool `json:"privacyprotectedallowed"`

	// Whether Order is Privacy Protected
	IsPrivacyProtected Bool `json:"isprivacyprotected"`

	// Whether Premium DNS is allowed for the Product Type
	PremiumDNSAllowed Bool `json:"premiumdnsallowed"`

	// Status of Premium DNS
	PremiumDNSEnabled Bool `json:"premiumdnsenabled"`

	// Whether Order Deletion is Allowed
	AllowDeletion Bool `json:"allowdeletion"`

	// Registrant Contact ID
	RegistrantContactID Int64 `json:"registrantcontactid"`

	// Registrant Contact Details
	RegistrantContact OrderContact `json:"registrantcontact"`

	// Admin Contact ID
	AdminContactID Int64 `json:"admincontactid"`

	// Admin Contact Details
	AdminContact OrderContact `json:"admincontact"`

	// Technical Contact ID
	TechContactID Int64 `json:"techcontactid"`

	// Technical Contact Details
	TechContact OrderContact `json:"techcontact"`

	// Billing Contact ID
	BillingContactID Int64 `json:"billingcontactid"`

	// Billing Contact Details
	BillingContact OrderContact `json:"billingcontact"`

	// Auto Renewal
	Recurring Bool `json:"recurring"`

	// Delegation Signer (DS) Record Details
	//  - Key Tag (keytag)
	//  - Algorithm (algorithm)
	//  - Digest Type (digesttype)
	//  - Digest (digest)
	DNSSec []string `json:"dnssec"`

	// GDPR Protection
	GDPR DomainGetOrderDetailsResponseGDPR `json:"gdpr"`

	Paused                   Bool    `json:"paused"`
	TNCRequired              Bool    `json:"tnc_required"`
	Actioncompleted          string  `json:"actioncompleted"`
	EntityID                 Int64   `json:"entityid"`
	ResellerCost             Float64 `json:"resellercost"`
	AutoRenewAttemptDuration Int64   `json:"autoRenewAttemptDuration"`
	AutoRenewTermType        string  `json:"autoRenewTermType"`
	ServiceProviderID        Int64   `json:"serviceproviderid"`
	MoneyBackPeriod          Int64   `json:"moneybackperiod"`
	EntityTypeID             Int64   `json:"entitytypeid"`
	ClassName                string  `json:"classname"`
	CustomerCost             Float64 `json:"customercost"`
	EaqID                    Int64   `json:"eaqid"`
	ClassKey                 string  `json:"classkey"`
	BulkWhoisOptOut          string  `json:"bulkwhoisoptout"`
	MultilingualFlag         string  `json:"multilingualflag"`
}

type DomainGetOrderDetailsResponseGDPR struct {
	Enabled  Bool `json:"enabled"`
	Eligible Bool `json:"eligible"`
}

type resDomainGetOrderDetailsResponse struct {
	errorResponse
	DomainGetOrderDetailsResponse
}

// GetOrderDetails Gets details of the Domain Registration Order associated with the specified Order Id.
// https://manage.resellerclub.com/kb/answer/770
func (domains *Domains) GetOrderDetails(orderID int64, options ...OrderDetailsOption) (*DomainGetOrderDetailsResponse, error) {
	u := domains.url("/details.json")
	q := u.Query()

	q.Set("order-id", strconv.FormatInt(orderID, 10))

	for _, opt := range options {
		q.Add("options", string(opt))
	}

	u.RawQuery = q.Encode()

	var res = resDomainGetOrderDetailsResponse{}
	err := domains.client.get(u.String(), &res)
	if err != nil {
		return nil, err
	}

	return &res.DomainGetOrderDetailsResponse, nil
}
