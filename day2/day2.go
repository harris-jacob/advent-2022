package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {

	file, err := os.Open("./input/day2.in")

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
	score := 0
	for scanner.Scan() {
		score += resultPartA(scanner.Text())
	}

	return score
}

func partTwo(scanner *bufio.Scanner) int {
	score := 0
	for scanner.Scan() {
		score += resultPartB(scanner.Text())
	}

	return score
}

func resultPartB(result string) int {
	switch result {
	case "A X":
		return 3
	case "A Y":
		return 4
	case "A Z":
		return 8
	case "B X":
		return 1
	case "B Y":
		return 5
	case "B Z":
		return 9
	case "C X":
		return 2
	case "C Y":
		return 6
	case "C Z":
		return 7
	}

	log.Fatal("unexpected input")
	return -1

}

func resultPartA(result string) int {
	switch result {
	case "A X":
		return 4
	case "A Y":
		return 8
	case "A Z":
		return 3
	case "B X":
		return 1
	case "B Y":
		return 5
	case "B Z":
		return 9
	case "C X":
		return 7
	case "C Y":
		return 2
	case "C Z":
		return 6
	}

	log.Fatal("unexpected input")
	return -1

}
