package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCrateLine(t *testing.T) {
	assert.Equal(t, []string{"", "D", ""}, parseCrateLine("    [D]    "))
	assert.Equal(t, []string{"N", "C", ""}, parseCrateLine("[N] [C]    "))
	assert.Equal(t, []string{"Z", "M", "P"}, parseCrateLine("[Z] [M] [P]"))
}

func TestParseCrates(t *testing.T) {
	assert.Equal(t,
		[][]string{
			{"N", "Z"},
			{"D", "C", "M"},
			{"P"},
		},
		parseCrates([]string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
		}),
	)
}

func TestParseMove(t *testing.T) {
	assert.Equal(t,
		move{count: 1, from: 2, to: 1},
		parseMove("move 1 from 2 to 1"),
	)
}

func TestMoveCrate(t *testing.T) {
	assert.Equal(t,
		[][]string{
			{"D", "N", "Z"},
			{"C", "M"},
			{"P"},
		},
		moveCrate(parseMove("move 1 from 2 to 1"), [][]string{
			{"N", "Z"},
			{"D", "C", "M"},
			{"P"},
		}),
	)

	assert.Equal(t,
		[][]string{
			{},
			{"C", "M"},
			{"Z", "N", "D", "P"},
		},
		moveCrate(parseMove("move 3 from 1 to 3"), [][]string{
			{"D", "N", "Z"},
			{"C", "M"},
			{"P"},
		}),
	)
}

func TestProcess(t *testing.T) {
	assert.Equal(t, "CMZ", Process("data_test.txt", moveCrate))
}

func TestMoveCrate1001(t *testing.T) {
	assert.Equal(t,
		[][]string{
			{"D", "N", "Z"},
			{"C", "M"},
			{"P"},
		},
		moveCrate1001(parseMove("move 1 from 2 to 1"), [][]string{
			{"N", "Z"},
			{"D", "C", "M"},
			{"P"},
		}),
	)

	assert.Equal(t,
		[][]string{
			{},
			{"C", "M"},
			{"D", "N", "Z", "P"},
		},
		moveCrate1001(parseMove("move 3 from 1 to 3"), [][]string{
			{"D", "N", "Z"},
			{"C", "M"},
			{"P"},
		}),
	)
}

func TestProcess1001(t *testing.T) {
	assert.Equal(t, "MCD", Process("data_test.txt", moveCrate1001))
}
