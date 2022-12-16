package main

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	println(Q1("data.txt"))
	println(Q2("data.txt"))
}

type Monkey struct {
	index        int
	items        []*big.Int
	operation    Operation
	test         Test
	inspectCount int
}

type Operation struct {
	leftOld  bool
	left     int
	rightOld bool
	right    int
	op       string
}

type Test struct {
	divisible int
	ifTrue    int
	ifFalse   int
}

func Process(input string) []*Monkey {
	lines := strings.Split(input, "\n\n")

	monkeys := make([]*Monkey, len(lines))

	for i, line := range lines {
		monkey := parseMonkey(line)
		monkeys[i] = &monkey
	}

	return monkeys
}

func parseMonkey(line string) Monkey {
	pattern, _ := regexp.Compile(
		`Monkey (\d+):
[ ]{2}Starting items: (.+)
[ ]{2}Operation: new = (old|\d+) ([*-/+]) (old|\d+)
[ ]{2}Test: divisible by (\d+)
[ ]{4}If true: throw to monkey (\d+)
[ ]{4}If false: throw to monkey (\d+)`)
	match := pattern.FindStringSubmatch(line)

	index, _ := strconv.Atoi(match[1])
	starting := strings.Split(match[2], ",")
	opLeft := match[3]
	op := match[4]
	opRight := match[5]
	divisible, _ := strconv.Atoi(match[6])
	ifTrue, _ := strconv.Atoi(match[7])
	ifFalse, _ := strconv.Atoi(match[8])

	intStarting := make([]*big.Int, len(starting))
	for i, s := range starting {
		intVal, _ := strconv.Atoi(strings.TrimSpace(s))
		intStarting[i] = big.NewInt(int64(intVal))
	}

	isLeftOld := opLeft == "old"
	isRightOld := opRight == "old"
	left := 0
	if !isLeftOld {
		left, _ = strconv.Atoi(opLeft)
	}
	right := 0
	if !isRightOld {
		right, _ = strconv.Atoi(opRight)
	}

	return Monkey{
		index,
		intStarting,
		Operation{isLeftOld, left, isRightOld, right, op},
		Test{divisible, ifTrue, ifFalse},
		0,
	}
}

func round(monkeys []*Monkey, bored bool) {
	totalDiv := monkeys[0].test.divisible
	for i := 1; i < len(monkeys); i++ {
		totalDiv *= monkeys[i].test.divisible
	}

	for _, monkey := range monkeys {
		fmt.Println(monkey.index, len(monkey.items))
		for _, w := range monkey.items {
			left := big.NewInt(int64(monkey.operation.left))
			if monkey.operation.leftOld {
				left = w
			}
			right := big.NewInt(int64(monkey.operation.right))
			if monkey.operation.rightOld {
				right = w
			}
			left = calc(left, right, monkey.operation.op)
			left = left.Mod(left, big.NewInt(int64(totalDiv)))

			if bored {
				left.Div(left, big.NewInt(3)) // bored
			}

			// fmt.Println(monkey.index, left, monkey.operation.op, right, new, monkey.test.divisible)

			mod := big.NewInt(0)
			mod.Mod(left, big.NewInt(int64(monkey.test.divisible)))
			if mod.Int64() == 0 {
				monkeys[monkey.test.ifTrue].items = append(monkeys[monkey.test.ifTrue].items, left)
			} else {
				monkeys[monkey.test.ifFalse].items = append(monkeys[monkey.test.ifFalse].items, left)
			}
			monkey.inspectCount++
		}
		monkey.items = monkey.items[:0]
	}
}

func calc(left *big.Int, right *big.Int, op string) *big.Int {
	switch op {
	case "+":
		return left.Add(left, right)
	case "-":
		return left.Sub(left, right)
	case "*":
		return left.Mul(left, right)
	case "/":
		return left.Div(left, right)
	}
	return left
}

func Q1(path string) int {
	raw, _ := os.ReadFile(path)
	monkeys := Process(string(raw))

	for i := 0; i < 20; i++ {
		fmt.Println("round", i)
		round(monkeys, true)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})

	fmt.Println(monkeys[0].inspectCount, monkeys[1].inspectCount)

	return monkeys[0].inspectCount * monkeys[1].inspectCount
}

func Q2(path string) int {
	raw, _ := os.ReadFile(path)
	monkeys := Process(string(raw))

	for i := 0; i < 10000; i++ {
		fmt.Println("round", i)
		round(monkeys, false)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})

	fmt.Println(monkeys[0].inspectCount, monkeys[1].inspectCount)

	return monkeys[0].inspectCount * monkeys[1].inspectCount
}
