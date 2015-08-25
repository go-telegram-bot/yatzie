package util

import (
	"math/rand"
)

func RandomFromArray(array []string) string {
	return array[rand.Intn(len(array))]
}

func StripPluginCommand(str string, prefix string, plugin string) string {
	return strings.Replace(str, prefix+plugin+" ", "", -1)
}
