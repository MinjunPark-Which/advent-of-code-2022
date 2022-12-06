package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	assert.Equal(t, 7, Detect("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
}
