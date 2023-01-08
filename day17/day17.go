package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day17.in")

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

	pattern := readPattern(scanner)
	rocks := make(map[[2]int]bool, 0)
	Ymax := 0
	i := 0

	for i := 0; i < 7; i++ {
		rocks[[2]int{i, 0}] = true
	}

	for rock := 0; rock < 2022; rock++ {
		fallingRock := getRockShape(rock%5, Ymax+4)

		for {
			// move sideways
			if pattern[i] == '>' {
				fallingRock = moveRight(fallingRock, rocks)
			} else {
				fallingRock = moveLeft(fallingRock, rocks)
			}

			i = (i + 1) % len(pattern)

			// move down
			var ok bool
			fallingRock, ok = moveDown(fallingRock, rocks)

			if !ok {
				// add to rocks
				for tile := range fallingRock {
					rocks[tile] = true
					if tile[1] > Ymax {
						Ymax = tile[1]
					}
				}
				break
			}

		}

	}

	return Ymax
}

func partTwo(scanner *bufio.Scanner) int {

	pattern := readPattern(scanner)
	rocks := make(map[[2]int]bool, 0)
	evaluated := make(map[[3]interface{}][2]int, 0)
	Ymax := 0
	extraHeight := 0
	i := 0

	for i := 0; i < 7; i++ {
		rocks[[2]int{i, 0}] = true
	}

	rock := 0
	for rock < 1000000000000 {
		fallingRock := getRockShape(rock%5, Ymax+4)

		for {
			// move sideways
			if pattern[i] == '>' {
				fallingRock = moveRight(fallingRock, rocks)
			} else {
				fallingRock = moveLeft(fallingRock, rocks)
			}

			i = (i + 1) % len(pattern)

			// move down
			var ok bool
			fallingRock, ok = moveDown(fallingRock, rocks)

			if !ok {

				// add to rocks
				for tile := range fallingRock {
					rocks[tile] = true
					if tile[1] > Ymax {
						Ymax = tile[1]
					}
				}

				state := [3]interface{}{rock % 5, i, stateSnapshot(rocks)}
				if deltas, ok := evaluated[state]; ok {
					dt := rock - deltas[0]
					dy := Ymax - deltas[1]


					cycles := (1000000000000 - rock) / dt

					extraHeight += cycles * dy
					rock += cycles * dt
					break
				}

				evaluated[[3]interface{}{rock % 5, i, stateSnapshot(rocks)}] = [2]int{rock, Ymax}
				break
			}
		}

		rock++

    }

	return Ymax + extraHeight
}

// get a state snapshot off the top 50 tiles
func stateSnapshot(rocks map[[2]int]bool) string {
	state := ""

	maxY := maxY(rocks)

	for y := 0; y < 50; y++ {
		for x := 0; x < 7; x++ {
			if rocks[[2]int{x, maxY - y}] {
				state += "#"
			} else {
				state += "."
			}
		}
	}

	return state
}

func maxY(rocks map[[2]int]bool) int {
	max := 0
	for tile := range rocks {
		if tile[1] > max {
			max = tile[1]
		}
	}
	return max
}

func moveLeft(rock map[[2]int]bool, rocks map[[2]int]bool) map[[2]int]bool {
	newRock := make(map[[2]int]bool)
	for tile := range rock {
		if tile[0] <= 0 || rocks[[2]int{tile[0] - 1, tile[1]}] {
			return rock
		}

		newRock[[2]int{tile[0] - 1, tile[1]}] = true
	}

	return newRock
}

func moveRight(rock map[[2]int]bool, rocks map[[2]int]bool) map[[2]int]bool {
	newRock := make(map[[2]int]bool)
	for tile := range rock {
		if tile[0] >= 6 || rocks[[2]int{tile[0] + 1, tile[1]}] {
			return rock
		}

		newRock[[2]int{tile[0] + 1, tile[1]}] = true
	}

	return newRock
}

func moveDown(rock map[[2]int]bool, rocks map[[2]int]bool) (map[[2]int]bool, bool) {
	newRock := make(map[[2]int]bool)
	for tile := range rock {
		if rocks[[2]int{tile[0], tile[1] - 1}] {
			return rock, false
		}

		newRock[[2]int{tile[0], tile[1] - 1}] = true
	}

	return newRock, true
}

func readPattern(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func printRocks(rock, fallingRock map[[2]int]bool, Ymax int) {
	for y := Ymax + 4; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if rock[[2]int{x, y}] {
				fmt.Print("#")
			} else if fallingRock[[2]int{x, y}] {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// ####

// .#.
// ###
// .#.

// ..#
// ..#
// ###

// #
// #
// #
// #

// ##
// ##
func getRockShape(rockNumber int, Y int) map[[2]int]bool {
	switch rockNumber {
	case 0:
		rock := make(map[[2]int]bool)
		rock[[2]int{2, Y}] = true
		rock[[2]int{3, Y}] = true
		rock[[2]int{4, Y}] = true
		rock[[2]int{5, Y}] = true

		return rock

	case 1:
		rock := make(map[[2]int]bool)
		rock[[2]int{3, Y + 2}] = true
		rock[[2]int{2, Y + 1}] = true
		rock[[2]int{3, Y + 1}] = true
		rock[[2]int{4, Y + 1}] = true
		rock[[2]int{3, Y}] = true

		return rock
	case 2:
		rock := make(map[[2]int]bool)
		rock[[2]int{4, Y + 2}] = true
		rock[[2]int{4, Y + 1}] = true
		rock[[2]int{4, Y}] = true
		rock[[2]int{3, Y}] = true
		rock[[2]int{2, Y}] = true

		return rock

	case 3:
		rock := make(map[[2]int]bool)
		rock[[2]int{2, Y + 3}] = true
		rock[[2]int{2, Y + 2}] = true
		rock[[2]int{2, Y + 1}] = true
		rock[[2]int{2, Y}] = true

		return rock

	default:
		rock := make(map[[2]int]bool)
		rock[[2]int{2, Y + 1}] = true
		rock[[2]int{3, Y + 1}] = true
		rock[[2]int{2, Y}] = true
		rock[[2]int{3, Y}] = true

		return rock

	}
}
