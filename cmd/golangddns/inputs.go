package main

import (
	"bufio"
	"log"
	"os"
)

func readInputAsString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Please input a valid string. ", err)
	}
	return input
}
