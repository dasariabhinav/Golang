package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/initializers"
	"project/network"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func init() {
	initializers.LoadEnvVariables()

}
func fetchDetails() map[string]interface{} {
	url := "https://api.ssllabs.com/api/v3/analyze?host=marriott.com"

	// Make a GET request to the API
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	// Unmarshal JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	return data
}
func validate() {
	data := fetchDetails()
	hostName, err := whois.Whois(data["host"].(string))

	hostStatus := data["status"]
	communicationProtocol := data["protocol"]

	//printing host status
	fmt.Println("Host status", hostStatus)

	//print communication protocol
	fmt.Println("Communication Protocol", communicationProtocol)

	if err == nil {
		result, e := whoisparser.Parse(hostName)
		if e == nil {
			// Print the domain status
			fmt.Println("Domain status", result.Domain.Status)
			// Print the domain created date
			fmt.Println("Creation Date", result.Domain.CreatedDate)
			// Print the domain expiration date
			fmt.Println("Expiration Date", result.Domain.ExpirationDate)
			// Print the registrar name
			fmt.Println("Registrar name", result.Registrar.Name)
			// Print the registrant name
			fmt.Println("Registrant name", result.Registrant.Name)
			// Print the registrant email address
			fmt.Println("email", result.Registrant.Email)
		}
	}
	network.Gethops()
}

func main() {
	validate()
}
