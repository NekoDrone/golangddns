package main

import (
	"log"
	"os"
)

func ipFileExists() bool {
	_, err := os.Stat(ipFilename)
	truth := err == nil
	if truth {
		log.Println("file exists")
	} else {
		log.Println("file doesn't exist.")
	}
	return truth
}

func readLastIpFromFile() string {
	log.Println("Reading last IP from file.")
	ipFile := openFile(ipFilename)
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

func createNewFile(filename string) {
	os.Create(filename)
}

func writeStringToFile(text string, filename string) {
	log.Printf("\nWriting string: %s to File", text)
	data := make([]byte, len(text))
	data = []byte(text)
	os.WriteFile(filename, data, 0666)
}
