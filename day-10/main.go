package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Q1("data.txt")
	Q2("data.txt")
}

func Q1(path string) int {
	result := 0
	cycles := Process(path)
	for i := 19; i < len(cycles); i += 40 {
		ss := (i + 1) * cycles[i-1]
		result += ss
		fmt.Println((i + 1), cycles[i-1], ss)
	}
	fmt.Println(result)
	return result
}

func Process(path string) []int {
	raw, _ := os.ReadFile(path)
	lines := strings.Split(string(raw), "\n")

	cycles := make([]int, 0)
	x := 1
	for _, line := range lines {
		command := strings.Split(line, " ")
		switch command[0] {
		case "noop":
			cycles = append(cycles, x)
		case "addx":
			cycles = append(cycles, x)
			y, _ := strconv.Atoi(command[1])
			x += y
			cycles = append(cycles, x)
		}
	}
	return cycles
}

func Q2(path string) string {
	cycles := Process(path)
	return strings.Join(render(cycles), "\n")
}

func render(cycles []int) []string {
	screen := make([]rune, 240)
	screen[0] = '#'
	for i := 1; i < len(screen); i++ {
		x := cycles[i-1]
		if within(i, x) {
			screen[i] = '#'
		} else {
			screen[i] = '.'
		}
	}

	output := make([]string, 6)
	for i := 0; i < 6; i++ {
		output[i] = string(screen[i*40 : i*40+40])
		fmt.Println(output[i])
	}
	return output
}

func within(i int, x int) bool {
	fmt.Println(i, x)
	return i%40 >= x-1 && i%40 <= x+1
}
