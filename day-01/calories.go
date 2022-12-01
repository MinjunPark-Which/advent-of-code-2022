package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func MaxCalory(data string) []int {
	caloryList := []int{}
	maxCalory := 0
	currentCalory := 0
	dataLines := strings.Split(data, "\n")
	for _, data := range dataLines {
		if calory, err := strconv.Atoi(data); err != nil {
			if maxCalory < currentCalory {
				maxCalory = currentCalory
			}
			caloryList = append(caloryList, currentCalory)
			currentCalory = 0
		} else {
			currentCalory += calory
		}
	}
	if maxCalory < currentCalory {
		maxCalory = currentCalory
	}
	if currentCalory > 0 {
		caloryList = append(caloryList, currentCalory)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloryList)))
	return caloryList
}

func main() {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	data := string(rawData)
	calories := MaxCalory(data)
	sum := calories[0] + calories[1] + calories[2]
	fmt.Print(sum)
}
