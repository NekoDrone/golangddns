package main

import (
	"bytes"
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
	recordToUpdate := readEnvVar("DNS_RECORD_NAME")

	body := *newDnsRecord(ip, recordToUpdate, "A")
	makeHttpRequest(http.MethodPut, apiUrl, body, apiKey)
}

type DNSRecord struct {
	Content string
	Name    string
	Type    string
}

func makeHttpRequest(method string, url string, body DNSRecord, authKey string) {
	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal("Could not convert struct to HTTP body. Something went wrong.", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(data))

	authToken := "Bearer " + authKey
	req.Header.Add("Authorization", authToken)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print("Could not send request to url. Check the internet connection? ", err)
	}

	log.Print(res)
}

func newDnsRecord(ipAddress string, dnsRecordName string, recordType string) *DNSRecord {
	return &DNSRecord{Content: ipAddress, Name: dnsRecordName, Type: recordType}
}
