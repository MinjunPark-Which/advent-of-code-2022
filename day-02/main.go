package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	data := string(rawData)
	win, shape := ScoreAll2(data)
	fmt.Printf("%v, %v: %v\n", shape, win, win+shape)
}

func getShapeScore(shape string) int {
	switch shape {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 0
}

func getWinScore(hand string) int {
	switch hand {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	return 0
}

func Score(left string, right string) (int, int) {
	winScore := 0

	leftShape := getShapeScore(left)
	rightShape := getShapeScore(right)

	if leftShape == rightShape {
		winScore = 3
	} else {
		if rightShape == 1 && leftShape == 3 {
			winScore = 6
		} else if rightShape == 2 && leftShape == 1 {
			winScore = 6
		} else if rightShape == 3 && leftShape == 2 {
			winScore = 6
		}
	}

	return winScore, rightShape
}

func ScoreAll(data string) (int, int) {
	winScore := 0
	shapeScore := 0

	dataLines := strings.Split(data, "\n")
	for _, line := range dataLines {
		if line == "" {
			continue
		}

		splitData := strings.Split(line, " ")
		win, shape := Score(splitData[0], splitData[1])
		winScore += win
		shapeScore += shape
	}

	return winScore, shapeScore
}

func Score2(left string, right string) (int, int) {
	rightShape := 0

	leftShape := getShapeScore(left)
	winScore := getWinScore(right)

	if winScore == 3 {
		rightShape = leftShape
	} else if winScore == 0 {
		switch leftShape {
		case 1:
			rightShape = 3
		case 2:
			rightShape = 1
		case 3:
			rightShape = 2
		}
	} else if winScore == 6 {
		switch leftShape {
		case 1:
			rightShape = 2
		case 2:
			rightShape = 3
		case 3:
			rightShape = 1
		}
	}

	return winScore, rightShape
}

func ScoreAll2(data string) (int, int) {
	winScore := 0
	shapeScore := 0

	dataLines := strings.Split(data, "\n")
	for i, line := range dataLines {
		if line == "" {
			continue
		}

		splitData := strings.Split(line, " ")
		win, shape := Score2(splitData[0], splitData[1])
		winScore += win
		shapeScore += shape
		fmt.Printf("%v: %v: %v %v\n", i, line, shape, win)
	}

	return winScore, shapeScore
}
