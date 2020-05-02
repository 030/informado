package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToByte(t *testing.T) {
	exp := "world"
	act := hello()
	assert.Equal(t, exp, act)
}
