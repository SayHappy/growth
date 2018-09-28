package helpers

import "math/rand"

func Random(min, max int) int {
	bettwen := max - min
	value := rand.Intn(bettwen) + min
	return value
}