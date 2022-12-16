package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ1(t *testing.T) {
	assert.Equal(t, 13140, Q1("data_test.txt"))
}

func TestQ2(t *testing.T) {
	assert.Equal(
		t,
		`##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`,
		Q2("data_test.txt"))
}
