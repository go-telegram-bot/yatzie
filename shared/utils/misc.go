package util

import (
	"math/rand"
)

func RandomFromArray(array []string) string {
	return array[rand.Intn(len(array))]
}
