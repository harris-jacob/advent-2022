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
		line := scanner.Text()
		fmt.Println(line)
		runes := []rune(line)

		for i := 0; i < len(runes)/2; i++ {
			for j := len(runes) - 1; j >= len(runes)/2; j-- {
				if runes[j] == runes[i] {
					fmt.Println("result", string(runes[i]))
					score += calcScore(int(runes[i]))
				}
			}
		}
	}

	return score

}

func partTwo(scanner *bufio.Scanner) string {
	return "NOT IMPLEMENTED"
}

func calcScore(asci int) int {
	if asci < 96 {
		return asci - 48
	} else {
		return asci - 96
	}
}
