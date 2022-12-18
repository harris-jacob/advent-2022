package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day9.in")

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

	H := []int{0, 0}
	T := []int{0, 0}
	visited := make(map[[2]int]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
		direction := direction(line[0])
		count, _ := strconv.Atoi(line[2:])

		for i := 0; i < count; i++ {
			H = moveHead(H, direction)
			T = moveTail(T, H)
			var key [2]int
			copy(key[:], T[:])
			visited[key] = true
		}

	}

	return len(visited)

}

func partTwo(scanner *bufio.Scanner) int {

	H := []int{0, 0}
	T := make([][]int, 9)

	for i := range T {
		T[i] = []int{0, 0}
	}

	visited := make(map[[2]int]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
		direction := direction(line[0])
		count, _ := strconv.Atoi(line[2:])

		for i := 0; i < count; i++ {
			H = moveHead(H, direction)
			T[0] = moveTail(T[0], H)
			for i := 1; i < len(T); i++ {
				T[i] = moveTail(T[i], T[i-1])
			}

			var key [2]int
			copy(key[:], T[8][:])
			visited[key] = true
		}

	}

	return len(visited)
}

func moveTail(tail []int, head []int) []int {
	dx := head[0] - tail[0]
	dy := head[1] - tail[1]

	if abs(dx) <= 1 && abs(dy) <= 1 {
		return tail
	}

	// part 2 motion types
	if abs(dx) >= 2 && abs(dy) >= 2 {
		x := 0
		y := 0
		if tail[0] < head[0] {
			x = head[0] - 1
		} else {
			x = head[0] + 1
		}

		if tail[1] < head[1] {
			y = head[1] - 1
		} else {
			y = head[1] + 1
		}

		return []int{x, y}

	}

	if abs(dx) >= 2 {
		if tail[0] < head[0] {
			return []int{head[0] - 1, head[1]}
		} else {
			return []int{head[0] + 1, head[1]}
		}
	} else if abs(dy) >= 2 {
		if tail[1] < head[1] {
			return []int{head[0], head[1] - 1}
		}
		return []int{head[0], head[1] + 1}
	}

	os.Exit(-1)
	return []int{}
}

func moveHead(head []int, direction []int) []int {
	x := head[0] + direction[0]
	y := head[1] + direction[1]

	return []int{x, y}
}

func direction(dir byte) []int {
	switch dir {
	case 'R':
		return []int{0, 1}

	case 'L':
		return []int{0, -1}

	case 'U':
		return []int{-1, 0}

	default:
		return []int{1, 0}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
