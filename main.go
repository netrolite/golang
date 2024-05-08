package main

import (
	"fmt"
)

const n int = 44345
const num = 3e-10

func main() {
	i := 0
	for {
		if i++; i > 20 {
			break
		}
		fmt.Printf("this is the fucking number: -----> %d <-----\n", i)
	}
}
