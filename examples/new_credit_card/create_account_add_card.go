package main

import (
	"github.com/hunterlong/authorizecim"
	"fmt"
	"os"
	"time"
	"math/rand"
)

func main() {

	apiName := os.Getenv("apiName")
	apiKey := os.Getenv("apiKey")
	AuthorizeCIM.SetAPIInfo(apiName,apiKey,"test")

	connected := AuthorizeCIM.TestConnection()

	if connected {
		fmt.Println("Successful Authorize.net Connection")
	} else {
		fmt.Println("There was an issue connecting to Authorize.net")
	}

	// Create random email address so it won't make duplicate records
	newUserEmail := RandomString(7)+"@domain.com"

	customer_info := AuthorizeCIM.AuthUser{
		"70",
		newUserEmail,
		"Test Account",
	}

	new_customer, success := AuthorizeCIM.CreateCustomerProfile(customer_info)

	if success {
		fmt.Println("New Customer Profile ID: ",new_customer)
	} else {
		fmt.Println("There was an issue creating the Customer Profile")
	}


	address := AuthorizeCIM.Address{
		FirstName: "Test",
		LastName: "User",
		Address: "1234 Road St",
		City: "City Name",
		State:" California",
		Zip: "93063",
		Country: "USA",
		PhoneNumber: "5555555555",
	}
	credit_card := AuthorizeCIM.CreditCard{
		CardNumber: "4111111111111111",
		ExpirationDate: "2020-12",
	}
	profile_id := new_customer
	newPaymentID, success := AuthorizeCIM.CreateCustomerBillingProfile(profile_id, credit_card, address)

	if success {
		fmt.Println("New Credit Card was added, Billing ID: ",newPaymentID)
	} else {
		fmt.Println("There was an issue inserting a credit card into the user account")
	}

}



// NOT NEEDED - ONLY FOR CREATING A RANDOM EMAIL ADDRESS
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}