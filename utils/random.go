package utils

import (
	"fmt"
	"math/rand"
)

func RandomChoice[T comparable](choices []T) T {
	return choices[rand.Intn(len(choices))]
}

func RandomShuffle[T comparable](slice []T) []T {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomPercentage() string {
	return fmt.Sprintf("%d%%", RandomNumber(0, 100))
}
