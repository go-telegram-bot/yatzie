package util

import (
	"math/rand"
	"strings"
)

func RandomFromArray(array []string) string {
	return array[rand.Intn(len(array))]
}

func StripPluginCommand(str string, prefix string, plugin string) string {
	return strings.Replace(str, prefix+plugin+" ", "", -1)
}
