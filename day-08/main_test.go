package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVisible(t *testing.T) {
	raw, _ := os.ReadFile("data_test.txt")
	trees := Parse(string(raw))
	assert.Equal(t, 21, CountVisible(trees))
}
