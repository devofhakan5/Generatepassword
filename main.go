package main

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const _charset = "abcdefghijklmnoABCDEFGHIJKLMNO123456789"

func GeneratePassword(length int) string {
	if length == 0 {
		return string("Number must be greater than 0")
	}

	x := make([]byte, length)
	for i := range x {
		x[i] = _charset[source.Int63()%int64(len(_charset))]
	}
	return string(x)
}
func main() {
	password := GeneratePassword(0)
	println(password)
}
