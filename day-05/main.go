package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Process("data.txt", moveCrate))
	fmt.Println(Process("data.txt", moveCrate1001))
}

type moveFunc func(m move, crates [][]string) [][]string

type move struct {
	count int
	from  int
	to    int
}

func parseCrateLine(line string) []string {
	stacks := make([]string, len(line)/3)
	i := 0
	for len(line) > i {
		crate := line[i : i+3]
		stacks[i/4] = strings.Trim(crate, " []")
		i += 4
	}

	return stacks
}

func parseCrates(lines []string) [][]string {
	stacks := make([][]string, (len(lines[0]))/3)
	for _, line := range lines {
		crates := parseCrateLine(line)
		for j, crate := range crates {
			if crate == "" {
				continue
			}
			stacks[j] = append(stacks[j], crate)
		}
	}
	return stacks
}

func parseMove(line string) move {
	pattern, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	match := pattern.FindStringSubmatch(line)
	if match == nil {
		return move{}
	}
	c, _ := strconv.Atoi(match[1])
	f, _ := strconv.Atoi(match[2])
	t, _ := strconv.Atoi(match[3])
	return move{count: c, from: f, to: t}
}

func moveCrate(m move, crates [][]string) [][]string {
	for i := 0; i < m.count; i++ {
		toMove := crates[m.from-1][0]
		crates[m.to-1] = append([]string{toMove}, crates[m.to-1]...) // move crate to new dest
		crates[m.from-1] = crates[m.from-1][1:]                      // delete crate from origin
	}
	return crates
}

func Process(path string, f moveFunc) string {
	raw, _ := os.ReadFile(path)
	data := string(raw)
	lines := strings.Split(data, "\n")
	// find split
	split := 0
	for i, line := range lines {
		if line == "" {
			split = i
			break
		}
	}

	crates := parseCrates(lines[:split-1])

	for _, m := range lines[split+1:] {
		crates = f(parseMove(m), crates)
	}

	output := ""
	for _, c := range crates {
		if len(c) > 0 {
			output += c[0]
		}
	}
	return output
}

func moveCrate1001(m move, crates [][]string) [][]string {
	toMove := make([]string, m.count)
	copy(toMove, crates[m.from-1][0:m.count])
	crates[m.from-1] = crates[m.from-1][m.count:]
	crates[m.to-1] = append(toMove, crates[m.to-1]...)
	return crates
}
