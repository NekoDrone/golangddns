package main

import (
	"log"
)

const ipFilename string = "lastIp"

func main() {
	log.Println("Starting DDNS Service.")
	var lastIp string

	if ipFileExists() {
		lastIp = readLastIpFromFile()
	} else {
		performFirstTimeSetup()
	}

	log.Println("Current Outgoing IP:")
	var currIp string = getCurrentOutgoingIpAsString()
	log.Println(currIp)

	if currIp != lastIp {
		writeStringToFile(currIp, ipFilename)
		updateIpOnExternalDnsRecord(currIp)
		log.Println("IP Address updated on external DNS site, and on local machine.")
	} else {
		log.Println("No IP changes have occurred in the last minute.")
	}
}

func performFirstTimeSetup() {
	createNewFile(ipFilename)
	createEnvFile()
	setVars()
}
