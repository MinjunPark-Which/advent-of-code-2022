package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	assert.Equal(t, 0, distance(Point{0, 0}, Point{0, 0}))
	assert.Equal(t, 0, distance(Point{1, 0}, Point{1, 0}))
	assert.Equal(t, 0, distance(Point{0, 1}, Point{0, 1}))
	assert.Equal(t, 0, distance(Point{1, 1}, Point{1, 1}))
	assert.Equal(t, 1, distance(Point{1, 0}, Point{0, 0}))
	assert.Equal(t, 1, distance(Point{0, 0}, Point{0, 1}))
	assert.Equal(t, 1, distance(Point{0, 0}, Point{1, 1}))
	assert.Equal(t, 2, distance(Point{0, 0}, Point{2, 2}))
	assert.Equal(t, 2, distance(Point{0, 0}, Point{2, 0}))
	assert.Equal(t, 3, distance(Point{0, 0}, Point{2, 1})) // actually incorrect
}

func TestQ1(t *testing.T) {
	assert.Equal(t, 13, Q1("data_test.txt"))
}

func TestQ2(t *testing.T) {
	assert.Equal(t, 1, Q2("data_test.txt"))
	assert.Equal(t, 36, Q2("data_test2.txt"))
}
