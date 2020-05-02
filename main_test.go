package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToByte(t *testing.T) {
	exp := "world"
	act, _ := hello()
	assert.Equal(t, exp, act)
}
