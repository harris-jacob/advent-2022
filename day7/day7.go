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
	file, err := os.Open("./input/day7.in")

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

	sizes := parse(scanner)
	sum := 0

	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func partTwo(scanner *bufio.Scanner) int {
	sizes := parse(scanner)

	currently_used := sizes["/"]
	allowed_used := 70_000_000 - 30_000_000
	need_to_free := currently_used - allowed_used

	min := 70_000_000
	for _, size := range sizes {
		if size >= need_to_free && size < min {
			min = size
		}
	}

	return min
}

func parse(scanner *bufio.Scanner) map[string]int {
	paths := make([]string, 0)
	sizes := make(map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if line[0] == "$" {
			switch line[1] {

			case "cd":
				cd(line[2], &paths)

			case "ls":
				continue
			}
		} else if line[0] == "dir" {
			continue
		} else {
			for i := 1; i <= len(paths); i++ {
				size, _ := strconv.Atoi(line[0])
				sizes[strings.Join(paths[:i], "/")] += size
			}
		}

	}

	return sizes
}

func cd(location string, paths *[]string) {
	if location == ".." {
		*paths = (*paths)[:len(*paths)-1]
	} else {
		*paths = append(*paths, location)
	}
}
