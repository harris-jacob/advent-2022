package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day1.in")

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
	maxCal := 0
	cal := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			num, _ := strconv.Atoi(text)
			cal += num
		} else {
			if cal > maxCal {
				maxCal = cal
			}

			cal = 0
		}
	}

	return maxCal
}

func partTwo(scanner *bufio.Scanner) int {
	elves := make([]int, 10)
	elf := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			num, _ := strconv.Atoi(text)
			elf += num
		} else {
			elves = append(elves, elf)
			elf = 0
		}
	}

	sort.Ints(elves)

	topThree := 0
	for i := 0; i < 3; i++ {
		topThree += elves[len(elves)-1-i]
	}

	return topThree
}
