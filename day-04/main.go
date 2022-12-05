package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Process("data.txt", isFullOverlap))
	fmt.Println(Process("data.txt", isIntersection))
}

type assignment struct {
	min int
	max int
}

type proc func([]assignment) bool

func parse(line string) []assignment {
	r := make([]assignment, 2)
	a := strings.Split(line, ",")
	for i, l := range a {
		s := strings.Split(l, "-")
		min, _ := strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])
		r[i] = assignment{min: min, max: max}
	}
	return r
}

func isFullOverlap(line []assignment) bool {
	if line[0].min <= line[1].min && line[0].max >= line[1].max {
		return true
	}
	if line[1].min <= line[0].min && line[1].max >= line[0].max {
		return true
	}
	return false
}

func Process(path string, f proc) int {
	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := string(rawData)
	lines := strings.Split(data, "\n")

	count := 0
	for _, line := range lines {
		if f(parse(line)) {
			count++
		}
	}

	return count
}

func isIntersection(line []assignment) bool {
	if line[0].min >= line[1].min && line[0].min <= line[1].max {
		return true
	} else if line[0].max >= line[1].min && line[0].max <= line[1].max {
		return true
	} else if line[1].min >= line[0].min && line[1].min <= line[0].max {
		return true
	} else if line[1].max >= line[0].min && line[1].max <= line[0].max {
		return true
	}
	return false
}
