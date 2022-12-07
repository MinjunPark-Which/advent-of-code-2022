package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func expected() Directory {
	expected := Directory{
		name:   "/",
		parent: nil,
		dirs:   make(map[string]*Directory),
		files: map[string]*File{
			"b.txt": {name: "b.txt", size: 14848514},
			"c.dat": {name: "c.dat", size: 8504156},
		},
		fileSize:  23352670,
		dirSize:   25028495,
		totalSize: 48381165,
	}

	a := &Directory{
		name:   "a",
		parent: &expected,
		dirs:   make(map[string]*Directory),
		files: map[string]*File{
			"f":     {name: "f", size: 29116},
			"g":     {name: "g", size: 2557},
			"h.lst": {name: "h.lst", size: 62596},
		},
		fileSize:  94269,
		dirSize:   584,
		totalSize: 94853,
	}
	expected.dirs["a"] = a

	e := &Directory{
		name:   "e",
		parent: a,
		dirs:   make(map[string]*Directory),
		files: map[string]*File{
			"i": {name: "i", size: 584},
		},
		fileSize:  584,
		dirSize:   0,
		totalSize: 584,
	}
	a.dirs["e"] = e

	d := &Directory{
		name:   "d",
		parent: &expected,
		dirs:   make(map[string]*Directory),
		files: map[string]*File{
			"j":     {name: "j", size: 4060174},
			"d.log": {name: "d.log", size: 8033020},
			"d.ext": {name: "d.ext", size: 5626152},
			"k":     {name: "k", size: 7214296},
		},
		fileSize:  24933642,
		dirSize:   0,
		totalSize: 24933642,
	}
	expected.dirs["d"] = d

	return expected
}

func TestParseLines(t *testing.T) {
	rawLines, _ := os.ReadFile("data_test.txt")
	lineArray := strings.Split(string(rawLines), "\n")
	root := ParseLines(lineArray)

	assert.Equal(t, expected(), root)
}

func TestTotalUnderLimit(t *testing.T) {
	rawLines, _ := os.ReadFile("data_test.txt")
	lineArray := strings.Split(string(rawLines), "\n")
	root := ParseLines(lineArray)

	start := int64(0)
	assert.Equal(t, int64(95437), TotalUnderLimit(&root, int64(100000), &start))
}

func TestDirsUnderLimit(t *testing.T) {
	//30000000
	rawLines, _ := os.ReadFile("data_test.txt")
	lineArray := strings.Split(string(rawLines), "\n")
	root := ParseLines(lineArray)

	out := make(map[string]int64)
	limit := 30000000 - (70000000 - root.totalSize)
	DirsOverLimit(&root, int64(limit), out)

	expected := map[string]int64{
		"d": 24933642,
		"/": 48381165,
	}

	assert.Equal(t, expected, out)
}

func TestMin(t *testing.T) {
	expected := map[string]int64{
		"d": 24933642,
		"/": 48381165,
	}

	assert.Equal(t, "d", Min(expected))
}
