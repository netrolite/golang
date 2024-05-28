package utils

import (
	"math/rand/v2"
)

func GetRandIntInRange(min int, max int) int {
	return rand.IntN(max-min) + min
}
