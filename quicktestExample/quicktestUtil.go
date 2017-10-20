package main

import (
	"math/rand"
)

func intBetween(rand *rand.Rand, min, max int) int {
	return rand.Intn(max-min) + min
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func stringWithLength(rand *rand.Rand, length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func oneOfColor(cs ...color) color {
	colors := []color{}
	for _, c := range cs {
		colors = append(colors, c)
	}
	return colors[rand.Intn(len(cs))]

}
