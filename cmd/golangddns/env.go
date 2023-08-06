package main

import (
	"log"
	"os"
	"strings"
)

func createEnvFile() {
	createNewFile(".env")
}

func setVars() {
	log.Print("Please enter your API Key here:")
	apiKey := readInputAsString()
	setApiKey(apiKey)

	log.Print("Please enter the API Endpoint you wish to update here:")
	apiEndpoint := readInputAsString()
	setApiEndpoint(apiEndpoint)

	log.Print("Please provide the name of the URL DNS Record you want to update (e.g. @ for root, and www for www) here:")
	dnsRecordName := readInputAsString()
	setDnsRecordName(dnsRecordName)

}

func setApiKey(value string) {
	apiKey := "API_KEY=" + value
	os.WriteFile(".env", []byte(apiKey), 0666) //this is not the proper way to do it, but it's limited so I"m leaving it be.
	//TODO: Write better .env accessor and maybe generalise the case.
	//or just use gotenv?
}

func setApiEndpoint(value string) {
	envFile := openFile(".env")
	currKeys := string(envFile[:])
	currKeys += "API_ENDPOINT=" + value
	os.WriteFile(".env", []byte(currKeys), 0666)
}

func setDnsRecordName(value string) {
	envFile := openFile(".env")
	currKeys := string(envFile[:])
	currKeys += "DNS_RECORD_NAME=" + value
	os.WriteFile(".env", []byte(currKeys), 0666)
}

func readEnvVar(envVar string) string {
	envFile := openFile(".env")
	allKeys := string(envFile[:])
	keyStrings := strings.Split(allKeys, "\n")
	for _, value := range keyStrings {
		temp := strings.Split(value, "=")
		if temp[0] == envVar {
			return temp[1]
		}
	}
	log.Fatalf("Unable to find given environment variable: %s", envVar)
	return ""
}
