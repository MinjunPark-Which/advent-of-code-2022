package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert.Equal(t, parse("2-4,6-8"),
		[]assignment{
			{min: 2, max: 4},
			{min: 6, max: 8},
		},
	)
}

func TestIsOverlap(t *testing.T) {
	assert.False(t, isFullOverlap(
		[]assignment{
			{min: 2, max: 4},
			{min: 6, max: 8},
		}),
	)

	assert.False(t, isFullOverlap(
		[]assignment{
			{min: 2, max: 3},
			{min: 4, max: 5},
		}),
	)

	assert.False(t, isFullOverlap(
		[]assignment{
			{min: 5, max: 7},
			{min: 7, max: 9},
		}),
	)

	assert.True(t, isFullOverlap(
		[]assignment{
			{min: 2, max: 8},
			{min: 3, max: 7},
		}),
	)

	assert.True(t, isFullOverlap(
		[]assignment{
			{min: 6, max: 6},
			{min: 4, max: 6},
		}),
	)

	assert.False(t, isFullOverlap(
		[]assignment{
			{min: 2, max: 6},
			{min: 4, max: 8},
		}),
	)
}

func TestProcess(t *testing.T) {
	assert.Equal(t,
		Process("data_test.txt", isFullOverlap),
		2,
	)
}

func TestIsIntersection(t *testing.T) {
	assert.False(t, isIntersection(
		[]assignment{
			{min: 2, max: 4},
			{min: 6, max: 8},
		}),
	)

	assert.False(t, isIntersection(
		[]assignment{
			{min: 2, max: 3},
			{min: 4, max: 5},
		}),
	)

	assert.True(t, isIntersection(
		[]assignment{
			{min: 5, max: 7},
			{min: 7, max: 9},
		}),
	)

	assert.True(t, isIntersection(
		[]assignment{
			{min: 2, max: 8},
			{min: 3, max: 7},
		}),
	)

	assert.True(t, isIntersection(
		[]assignment{
			{min: 6, max: 6},
			{min: 4, max: 6},
		}),
	)

	assert.True(t, isIntersection(
		[]assignment{
			{min: 2, max: 6},
			{min: 4, max: 8},
		}),
	)

}
