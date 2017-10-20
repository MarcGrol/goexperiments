package main

import (
	"testing"
	"testing/quick"
)

func TestValidatePerson(t *testing.T) {
	assertion := func(p person) bool {
		return p.IsCorrect()
	}

	// Run the assertion through the quick checker
	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}
