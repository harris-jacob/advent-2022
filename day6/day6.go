package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day6.in")

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
	return findMarker(scanner, 4)
}

func partTwo(scanner *bufio.Scanner) int {
	return findMarker(scanner, 14)
}

func findMarker(scanner *bufio.Scanner, n int) int {
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)-n; i++ {
			if unique(line[i : i+n]) {
				return i + n
			}
		}
	}

	return -1
}

func unique(chars string) bool {
	set := make(map[rune]bool)

	for _, r := range chars {
		set[r] = true
	}

	return len(chars) == len(set)

}
