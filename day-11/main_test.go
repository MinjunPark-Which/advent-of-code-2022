package main

import (
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ1(t *testing.T) {
	result := Q1("data_test.txt")
	assert.Equal(t, 10605, result)
}

func TestQ2(t *testing.T) {
	result := Q2("data_test.txt")
	assert.Equal(t, 2713310158, result)
}

func TestParseMonkey(t *testing.T) {
	raw, _ := os.ReadFile("data_test.txt")
	lines := strings.Split(string(raw), "\n\n")

	monkey := parseMonkey(lines[0])

	assert.Equal(t,
		Monkey{
			0,
			[]*big.Int{big.NewInt(int64(79)), big.NewInt(int64(98))},
			Operation{true, 0, false, 19, "*"},
			Test{23, 2, 3},
			0,
		},
		monkey,
	)

	monkey = parseMonkey(lines[3])

	assert.Equal(t,
		Monkey{
			3,
			[]*big.Int{big.NewInt(int64(74))},
			Operation{true, 0, false, 3, "+"},
			Test{17, 0, 1},
			0,
		},
		monkey,
	)
}
