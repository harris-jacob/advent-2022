package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day4.in")

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
	count := 0

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")

		range1 := strings.Split(pairs[0], "-")
		range2 := strings.Split(pairs[1], "-")

		min1, _ := strconv.ParseInt(range1[0], 10, 64)
		max1, _ := strconv.ParseInt(range1[1], 10, 64)
		min2, _ := strconv.ParseInt(range2[0], 10, 64)
		max2, _ := strconv.ParseInt(range2[1], 10, 64)

		if min1 >= min2 && max1 <= max2 {
			count++
		} else if min2 >= min1 && max2 <= max1 {
			count++
		}

	}
	return count
}

func partTwo(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")

		range1 := strings.Split(pairs[0], "-")
		range2 := strings.Split(pairs[1], "-")

		min1, _ := strconv.ParseInt(range1[0], 10, 64)
		max1, _ := strconv.ParseInt(range1[1], 10, 64)
		min2, _ := strconv.ParseInt(range2[0], 10, 64)
		max2, _ := strconv.ParseInt(range2[1], 10, 64)

		if min1 >= min2 && min1 <= max2 {
			count++
		} else if max1 >= min2 && max1 <= max2 {
			count++
		} else if min2 >= min1 && min2 <= max1 {
			count++
		} else if max2 >= min1 && max2 <= max1 {
			count++
		}

	}

	return count
}
