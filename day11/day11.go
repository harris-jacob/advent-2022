package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day11.in")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := utils.ScannerFromStart(file)

	fmt.Println(partOne(scanner))

	scanner = utils.ScannerFromStart(file)

	fmt.Println(partTwo(scanner))
}

func partOne(scanner *bufio.Scanner) int {
	monkeys := make([]*Monkey, 0)

	for scanner.Scan() {
		monkeys = append(monkeys, parseMonkey(scanner))
	}

    // 20 turns
	for i := 0; i < 20; i++ {
		turn(monkeys)
	}

    sort.Slice(monkeys, func(i, j int) bool {
        return monkeys[i].inspectCount > monkeys[j].inspectCount
    })



	return monkeys[0].inspectCount * monkeys[1].inspectCount
}

func partTwo(scanner *bufio.Scanner) int {
	monkeys := make([]*Monkey, 0)

	for scanner.Scan() {
		monkeys = append(monkeys, parseMonkey(scanner))
	}

    lcm := 1
    for _, monkey := range monkeys {
        lcm = lcm * monkey.divisibleBy
    }

	for i := 0; i < 10000; i++ {
		turn(monkeys, lcm)
	}

    sort.Slice(monkeys, func(i, j int) bool {
        return monkeys[i].inspectCount > monkeys[j].inspectCount
    })



	return monkeys[0].inspectCount * monkeys[1].inspectCount
}

type Monkey struct {
	items       []int
	operation   string
	divisibleBy int
	trueThen    int
	falseThen   int
    inspectCount int
}

type Move struct {
	item int
	to   int
}

func turn(monkeys []*Monkey, args ...int) {
	for _, monkey := range monkeys {
		toMove := make([]*Move, 0)
		for idx := range monkey.items {
			var move Move
			inspectItem(monkey, idx)
			// monkey gets bored
            if len(args) == 0 {
			    monkey.items[idx] /= 3
            } else {
                monkey.items[idx] %= args[0]
            }
			
            move.item = monkey.items[idx]
			
            if monkey.items[idx]%monkey.divisibleBy == 0 {
				move.to = monkey.trueThen
			} else {
				move.to = monkey.falseThen
			}

			toMove = append(toMove, &move)
		}

		// monkey has nothing left
		monkey.items = make([]int, 0)

		// perform moves
		for _, move := range toMove {
			monkeys[move.to].items = append(monkeys[move.to].items, move.item)
		}
	}
}

func inspectItem(monkey *Monkey, idx int) {
	if monkey.operation[12:] == "old" {
		monkey.items[idx] = performOperation(monkey.operation[10], monkey.items[idx], monkey.items[idx])
	} else {
		val, _ := strconv.Atoi(monkey.operation[12:])
		monkey.items[idx] = performOperation(monkey.operation[10], monkey.items[idx], val)
	}

    monkey.inspectCount++

}

func performOperation(op byte, old, value int) int {
	switch rune(op) {
	case '+':
		return old + value
	case '-':
		return old - value
	default:
		return old * value
	}
}

func parseMonkey(scanner *bufio.Scanner) *Monkey {
	var monkey Monkey

	scanner.Scan()
	monkey.items = parseItems(scanner.Text())
	scanner.Scan()
	monkey.operation = parseOperation(scanner.Text())
	scanner.Scan()
	monkey.divisibleBy = parseDivisibleBy(scanner.Text())
	scanner.Scan()
	monkey.trueThen = parseBoolCond(scanner.Text())
	scanner.Scan()
	monkey.falseThen = parseBoolCond(scanner.Text())
	scanner.Scan()

	return &monkey

}

func parseNumber(line string) int {
	fmt.Println(line)
	split := strings.Split(line, " ")[1]
	count, _ := strconv.Atoi(split[:len(split)-1])

	return count
}

func parseItems(line string) []int {
	items := make([]int, 0)
	split := strings.Split(line, ":")[1]

	stringItems := strings.Split(split, ",")

	for _, item := range stringItems {
		val, _ := strconv.Atoi(strings.Trim(item, " "))
		items = append(items, val)
	}

	return items

}

func parseOperation(line string) string {
	return strings.Trim(strings.Split(line, ":")[1], " ")
}

func parseDivisibleBy(line string) int {
	result, _ := strconv.Atoi(strings.Trim(strings.Split(line, ":")[1], " ")[13:])

	return result
}

func parseBoolCond(line string) int {
	result, _ := strconv.Atoi(strings.Trim(strings.Split(line, ":")[1], " ")[16:])
	return result
}
