package utils

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func ReadString(rd *bufio.Reader) string {
	val, err := rd.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	val = strings.TrimSpace(val)

	return val
}

func ReadInt(rd *bufio.Reader) int {
	valStr, err := rd.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	valStr = strings.TrimSpace(valStr)

	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Fatal(err)
	}

	return val
}
