package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Ltime)

	msg, err := greet("John")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(msg)
	}
}

func greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("name must not be empty")
	}
	return fmt.Sprintf("Hello, %v!", name), nil
}
