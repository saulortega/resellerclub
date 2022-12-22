# resellerclub

## ⚠ Supports only a few features. Development ongoing.

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/saulortega/resellerclub"
)

func main() {
	// Create Client:
	client := resellerclub.New("123456", "asdfghjklqwertyuiopzxcvbnm")

	// Check domains availability:
	domains := []string{"example", "ejemplo"}
	tlds := []string{"com", "co"}
	domainsAvalilability, err := client.Domains.CheckAvailability(domains, tlds)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" -> Domains.CheckAvailability:: %+v\n", domainsAvalilability)

	// Get the Order ID of a registered domain name:
	orderID, err := client.Domains.GetOrderID("example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" -> Domains.GetOrderID:: %v\n", orderID)

	// Get details of the Domain Registration Order associated with the specified Order ID:
	res, err := client.Domains.GetOrderDetails(123456, resellerclub.OrderDetailsOptionAll)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" -> Domains.GetOrderDetails:: %+v\n", res)

	// Get a list of Domain Registration Orders matching the search criteria, along with the details:
	params := &resellerclub.DomainSearchParams{
		NoOfRecords: 10,
		PageNo:      1,
		CustomerIDs: []int64{123456},
	}
	res, err := client.Domains.Search(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" -> Domains.Search:: %+v\n", res)

	// Create a new customer:
	params := &resellerclub.CustomerCreateParams{
		Username:     "me@example.com",
		Password:     "Abc123.!%",
		Name:         "Jhon",
		Company:      "MyCompany",
		AddressLine1: "Addr1",
		City:         "City1",
		State:        "St",
		Country:      "CO",
		Zipcode:      "000000",
		PhoneCC:      "57",
		Phone:        "1234567890",
		LangPref:     "es",
	}
	customerID, err := client.Customers.Create(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" -> Customers.Create:: %v\n", customerID)
}
```
