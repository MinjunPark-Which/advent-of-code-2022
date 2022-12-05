package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	left, right := split("abcABC")
	assert.Equal(t, left, "abc")
	assert.Equal(t, right, "ABC")
}

func TestToPriority(t *testing.T) {
	assert.Equal(t, toPriority('a'), 1)
	assert.Equal(t, toPriority('c'), 3)
	assert.Equal(t, toPriority('C'), 29)
}

func TestGetPriority(t *testing.T) {
	assert.Equal(t, GetPriority("vJrwpWtwJgWrhcsFMMfFFhFp"), 16)
	assert.Equal(t, GetPriority("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), 38)
	assert.Equal(t, GetPriority("PmmdzqPrVvPwwTWBwg"), 42)
	assert.Equal(t, GetPriority("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"), 22)
	assert.Equal(t, GetPriority("ttgJtRGJQctTZtZT"), 20)
	assert.Equal(t, GetPriority("CrZsJsPPZsGzwwsLwLmpwMDw"), 19)
}

func TestProcess(t *testing.T) {
	assert.Equal(t, Process("data_test.txt"), 157)
}

func TestFindGroupBadge(t *testing.T) {
	assert.Equal(t,
		findGroupBadge([]string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
		}),
		18,
	)

	assert.Equal(t,
		findGroupBadge([]string{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}),
		52,
	)
}

func TestProcessGroup(t *testing.T) {
	assert.Equal(t, ProcessGroupBadge("data_test.txt"), 70)
}
