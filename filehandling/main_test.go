package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFact(t *testing.T) {
	expectedVal := "This needs to go in a file - LearnCodeOnline.in"
	actualVal := string(databyte)
	assert.Equal(t, expectedVal, actualVal, "value is not correct")
}
