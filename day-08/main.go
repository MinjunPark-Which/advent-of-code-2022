package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	raw, _ := os.ReadFile("data.txt")
	trees := Parse(string(raw))
	// fmt.Println(CountVisible(trees))
	fmt.Println(ScenicScore(trees))
}

type tree struct {
	x      int
	y      int
	height int
	left   *tree
	down   *tree
	up     *tree
	right  *tree
}

func Parse(data string) [][]tree {
	var lines = strings.Split(data, "\n")
	var matrix = make([][]int, len(lines))

	for i, line := range lines {
		for j, r := range line {
			if j == 0 {
				matrix[i] = make([]int, len(line))
			}
			matrix[i][j] = int(r - '0')
		}
	}

	var trees = make([][]tree, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				trees[i] = make([]tree, len(matrix[i]))
			}
			trees[i][j] = tree{height: matrix[i][j], x: i, y: j}
		}
	}

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			populateSurrounds(i, j, trees)
		}
	}

	return trees
}

func populateSurrounds(x int, y int, trees [][]tree) {
	if x > 0 {
		trees[x][y].up = &trees[x-1][y]
	}
	if x < len(trees)-1 {
		trees[x][y].down = &trees[x+1][y]
	}
	if y > 0 {
		trees[x][y].left = &trees[x][y-1]
	}
	if y < len(trees)-1 {
		trees[x][y].right = &trees[x][y+1]
	}
}

func CountVisible(trees [][]tree) int {
	var count = 0
	for _, line := range trees {
		for _, point := range line {
			if isVisible(point, func(point tree) *tree { return point.up }) ||
				isVisible(point, func(point tree) *tree { return point.right }) ||
				isVisible(point, func(point tree) *tree { return point.down }) ||
				isVisible(point, func(point tree) *tree { return point.left }) {
				count++
			}
		}
	}
	return count
}

type fetchAdjacent func(tree) *tree

func fetchLineOfSight(point tree, fn fetchAdjacent) []tree {
	var lineOfSight = make([]tree, 0)
	var adjacent = fn(point)
	for adjacent != nil {
		lineOfSight = append(lineOfSight, *adjacent)
		adjacent = fn(*adjacent)
	}
	return lineOfSight
}

func isVisible(point tree, fn fetchAdjacent) bool {
	var lineOfSight = fetchLineOfSight(point, fn)

	//debug
	var lines = make([]int, 0)
	for _, line := range lineOfSight {
		lines = append(lines, line.height)
	}
	fmt.Println(point.x, point.y, point.height, lines)
	// end debug

	if len(lineOfSight) == 0 {
		return true
	}
	for _, line := range lineOfSight {
		if line.height >= point.height {
			return false
		}
	}
	return true
}

func ScenicScore(trees [][]tree) int {
	var score = make([][]int, len(trees))
	for i, line := range trees {
		score[i] = make([]int, len(line))
		for j, point := range line {
			var up = fetchLineOfSight(point, func(point tree) *tree { return point.up })
			var right = fetchLineOfSight(point, func(point tree) *tree { return point.right })
			var down = fetchLineOfSight(point, func(point tree) *tree { return point.down })
			var left = fetchLineOfSight(point, func(point tree) *tree { return point.left })
			score[i][j] = fetchScenicScore(point, up) *
				fetchScenicScore(point, right) *
				fetchScenicScore(point, down) *
				fetchScenicScore(point, left)
		}
	}

	var highestScore = 0
	for _, a := range score {
		for _, b := range a {
			if b > highestScore {
				highestScore = b
			}
		}
	}
	return highestScore
}

func fetchScenicScore(point tree, lineOfSight []tree) int {
	var i = 0
	for _, sight := range lineOfSight {
		i++
		if sight.height >= point.height {
			return i
		}
	}
	return i
}
