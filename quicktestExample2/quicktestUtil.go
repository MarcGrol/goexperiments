package main

import (
	"math/rand"
)

func randomInt(rand *rand.Rand, min, max int) int {
	return rand.Intn(max-min) + min
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(rand *rand.Rand, length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
