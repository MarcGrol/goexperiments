package main

import (
	"log"
	"math/rand"
	"reflect"
)

type person struct {
	name string
	age  int
}

func (p person) Generate(rand *rand.Rand, size int) reflect.Value {
	randomP := person{
		name: isStringWithLength(rand, isIntBetween(rand, 1, 32)),
		age:  isIntBetween(rand, 0, 100),
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

	return true
}
