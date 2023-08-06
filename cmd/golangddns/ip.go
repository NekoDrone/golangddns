package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type IP struct {
	Query string
}

func getCurrentOutgoingIpAsString() string {
	req := requestApiToGetIpData("http://ip-api.com/json/")
	defer (*req).Body.Close()
	data := convertResponseBodyToByteArray(req)
	return unmarshalJsonByteArrayToIpString(data)
}

func requestApiToGetIpData(url string) *http.Response {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to get request from external ip-api.\nIs the service down?", err)
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

func unmarshalJsonByteArrayToIpString(byteArray []byte) string {
	var ip IP
	json.Unmarshal(byteArray, &ip)
	return ip.Query
}

func updateIpOnExternalDnsRecord(ip string) {
	apiKey := readEnvVar("API_KEY")
	apiUrl := readEnvVar("API_ENDPOINT")

	log.Println("This is where I would have written the API request.")
	log.Println("The outgoing update would have contained this IP:")
	log.Println(ip)
	log.Println("We would have sent the request to the following endpoint:")
	log.Println(apiUrl)
	log.Println("API Key:")
	log.Println(apiKey)
}
