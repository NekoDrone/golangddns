package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const ipFile string = "lastIp"

func main() {
	log.Println("Starting DDNS Service.")
	var lastIp string

	if ipFileExists() {
		lastIp = readLastIpFromFile()
	} else {
		createLastIpFile()
	}

	log.Println("\nCurrent Outgoing IP:")
	var currIp string = getCurrentOutgoingIpAsString()
	log.Println(currIp)

	if currIp != lastIp {
		writeIpToFile(currIp)
		updateIpOnExternalDnsRecord(currIp)
		log.Println("IP Address updated on external DNS site, and on local machine.")
	} else {
		log.Println("No IP changes have occurred in the last minute.")
	}
}

func getCurrentOutgoingIpAsString() string {
	req := requestApiToGetIpData("http://ip-api.com/json/")
	defer (*req).Body.Close()
	data := convertResponseBodyToByteArray(req)
	return unmarshalByteArrayToIpString(data)
}

func requestApiToGetIpData(url string) *http.Response {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to get request from external API.\n", err)
	}
	return req
}

func convertResponseBodyToByteArray(responsePointer *http.Response) []byte {

	response := *responsePointer
	/*
		http.Get returns a pointer to the response, so we resolve it here.
		Technically, accessing a field or property on a pointer automatically
		resolves the pointer, but that's insane behaviour and I'm resolving it for my own sanity.
	*/
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Response body converted to invalid array.", err)
	}
	return data
}

func unmarshalByteArrayToIpString(byteArray []byte) string {
	var ip IP
	json.Unmarshal(byteArray, &ip)
	return ip.Query
}

type IP struct {
	Query string
}

func readLastIpFromFile() string {
	log.Println("Reading last IP from file.")
	ipFile := openFile(ipFile)
	res := string(ipFile[:])
	log.Println("Last IP found at:")
	log.Println(res)
	return res
}

func openFile(filename string) []byte {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not open file 'lastIP'. Please check access rights.")
	}
	return fileData
}

func createLastIpFile() {
	os.Create(ipFile)
}

func ipFileExists() bool {
	_, err := os.Stat(ipFile)
	truth := err == nil
	if truth {
		log.Println("file exists")
	} else {
		log.Println("file doesn't exist.")
	}
	return truth
}

func writeIpToFile(ip string) {
	log.Println("Writing IP to File")
	data := make([]byte, len(ip))
	os.WriteFile(ipFile, data, 0666)
}

func updateIpOnExternalDnsRecord(ip string) {
	log.Println("This is where I would have written the API request.")
	log.Println("The outgoing update would have contained this IP:")
	log.Println(ip)
}
