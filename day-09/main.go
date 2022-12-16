package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Q1("data.txt"))
	fmt.Println(Q2("data.txt"))
}

type Point struct {
	x, y int
}

func Q1(path string) int {
	raw, _ := os.ReadFile(path)
	tailVisits := Process(string(raw), 2)
	return uniqueCount(tailVisits)
}

func Q2(path string) int {
	raw, _ := os.ReadFile(path)
	tailVisits := Process(string(raw), 10)
	return uniqueCount(tailVisits)
}

func Process(input string, size int) []Point {
	tails := make([]Point, size)
	for i := 0; i < len(tails); i++ {
		tails[i] = Point{0, 0}
	}
	tailVisits := make([]Point, 0)
	tailVisits = append(tailVisits, Point{0, 0})

	i := 0
	for _, l := range strings.Split(input, "\n") {
		newCoords := parseLine(l)
		for _, c := range newCoords {
			var newHead Point

			// tail movements
			for j := 0; j < len(tails); j++ {
				if j == 0 {
					newHead = add(tails[j], c)
				} else {
					newHead = moveTail(newHead, tails[j])
				}
				tails[j] = newHead

				if j == len(tails)-1 {
					tailVisits = append(tailVisits, tails[j])
				}
			}

			fmt.Println(l, tails)
			i++
		}
	}

	return tailVisits
}

func moveTail(head Point, tail Point) Point {
	x := head.x - tail.x
	y := head.y - tail.y

	newTail := Point{tail.x, tail.y}

	if abs(x) > 1 && abs(y) > 1 {
		newTail.x += (x / abs(x))
		newTail.y += (y / abs(y))
	} else if abs(x) > 1 {
		newTail.x += (x / abs(x))
		newTail.y = head.y
	} else if abs(y) > 1 {
		newTail.x = head.x
		newTail.y += (y / abs(y))
	}

	return newTail
}

func distance(head Point, tail Point) int {
	x := abs(head.x - tail.x)
	y := abs(head.y - tail.y)

	if x == 1 && y == 1 {
		return 1
	}

	return x + y
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func parseLine(line string) []Point {
	command := strings.Split(line, " ")
	dist, _ := strconv.Atoi(command[1])
	result := make([]Point, dist)

	for i := 0; i < dist; i++ {
		switch command[0] {
		case "R":
			result[i] = Point{1, 0}
		case "U":
			result[i] = Point{0, 1}
		case "D":
			result[i] = Point{0, -1}
		case "L":
			result[i] = Point{-1, 0}
		}
	}

	return result
}

func add(source Point, dist Point) Point {
	source.x = source.x + dist.x
	source.y = source.y + dist.y
	return source
}

func uniqueCount(visits []Point) int {
	uniques := make(map[Point]bool, 0)
	for _, v := range visits {
		uniques[v] = true
	}
	return len(uniques)
}
