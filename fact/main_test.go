package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFact(t *testing.T) {
	expectedVal := 120.0
	actualVal := fact(5)
	assert.Equal(t, expectedVal, actualVal, "value is not correct")
}
