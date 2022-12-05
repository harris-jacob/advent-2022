package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day3.in")

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
	// iterate lines
	for scanner.Scan() {
		line := []rune(scanner.Text())

		set := make(map[rune]bool)

		for i := 0; i < len(line)/2; i++ {
			set[line[i]] = true
		}

		for i := len(line) / 2; i < len(line); i++ {
			if set[line[i]] {
				score += calcScore(int(line[i]))
				break
			}
		}
	}

	return score

}

func partTwo(scanner *bufio.Scanner) int {
	score := 0
	lines := make([]string, 0)

	// read to array
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 2; i < len(lines); i += 3 {

		// set of line 1 runes
		line1Set := make(map[rune]bool)
		for _, val := range lines[i] {
			line1Set[val] = true
		}

		// intersection of 1 and 2
		temp := make(map[rune]bool)
		for _, val := range lines[i-1] {
			if line1Set[val] {
				temp[val] = true
			}
		}

		// intersection of temp and 3
		for _, val := range lines[i-2] {
			if temp[val] {
				score += calcScore(int(val))
				break
			}
		}
	}

	return score

}

func calcScore(asci int) int {
	if asci < 96 {
		return asci - 38
	} else {
		return asci - 96
	}
}
