package main

import (
	"log"
	"math/rand"
	"reflect"
)

type color int

const (
	red color = iota
	green
	blue
	black
	white
)

type person struct {
	name          string
	age           int
	favoriteColor color
}

func (p person) Generate(rand *rand.Rand, size int) reflect.Value {
	randomP := person{
		name:          stringWithLength(rand, intBetween(rand, 1, 32)),
		age:           intBetween(rand, 0, 100),
		favoriteColor: oneOfColor(red, green, blue),
	}
	log.Printf("person: %#v", randomP)
	return reflect.ValueOf(randomP)
}

func (p person) IsCorrect() bool {
	if len(p.name) < 1 || len(p.name) > 32 {
		return false
	}

	if p.age < 0 || p.age > 100 {
		return false
	}

	if p.favoriteColor == black || p.favoriteColor == white {
		return false
	}

	return true
}
