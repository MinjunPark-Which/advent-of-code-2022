package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	result := Process("data.txt")
	fmt.Println(result)

	result = ProcessGroupBadge("data.txt")
	fmt.Println(result)
}

func split(items string) (string, string) {
	middle := len(items) / 2
	return items[:middle], items[middle:]
}

func toPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item-'A') + 27
	}

	return int(item-'a') + 1
}

func GetPriority(items string) int {
	left, right := split(items)
	m := make(map[rune]bool)
	p := 0
	for _, c := range left {
		m[c] = true
	}
	for _, c := range right {
		if m[c] {
			p = toPriority(c)
			break
		}
	}
	return p
}

func Process(path string) int {
	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := string(rawData)
	lines := strings.Split(data, "\n")

	sum := 0
	for _, line := range lines {
		sum += GetPriority(line)
	}
	return sum
}

func findGroupBadge(group []string) int {
	x := make(map[rune]bool)
	for _, a := range group[0] {
		x[a] = true
	}

	y := make(map[rune]bool)
	for _, b := range group[1] {
		if x[b] {
			y[b] = true
		}
	}

	for _, c := range group[2] {
		if y[c] {
			return toPriority(c)
		}
	}

	return 0
}

func ProcessGroupBadge(path string) int {
	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := string(rawData)
	lines := strings.Split(data, "\n")

	sum := 0
	for i := 0; i < len(lines)/3; i++ {
		sum += findGroupBadge(lines[i*3 : i*3+3])
	}
	return sum
}
