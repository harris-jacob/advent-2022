package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day8.in")

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
	text := make([][]string, 0)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	i := 0
	for scanner.Scan() {
		text = append(text, make([]string, 0))
		for _, v := range scanner.Text() {
			text[i] = append(text[i], string(v))
		}
		i++
	}

	visible := 0
	for row := 0; row < len(text); row++ {
		for column := 0; column < len(text[1]); column++ {
			for _, direction := range directions {
				nextRow := row
				nextCol := column
				vis := true

				for {
					nextRow += direction[0]
					nextCol += direction[1]
					if (nextRow >= 0 && nextRow < len(text)) && (nextCol >= 0 && nextCol < len(text[0])) {
						if text[nextRow][nextCol] >= text[row][column] {
							vis = false
						}

					} else {
						break
					}
				}

				if vis {
					visible++
					break
				}
			}
		}
	}

	return visible
}

func partTwo(scanner *bufio.Scanner) int {
	text := make([][]string, 0)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	i := 0
	for scanner.Scan() {
		text = append(text, make([]string, 0))
		for _, v := range scanner.Text() {
			text[i] = append(text[i], string(v))
		}
		i++
	}

	max := 0
	for row := 0; row < len(text); row++ {
		for column := 0; column < len(text[1]); column++ {
			score := 1
			for _, direction := range directions {
				dirDist := 0
				nextRow := row
				nextCol := column

				for {
					nextRow += direction[0]
					nextCol += direction[1]

					dirDist++
					// if we aren't at an edge
					if (nextRow >= 0 && nextRow < len(text)) && (nextCol >= 0 && nextCol < len(text[0])) {
						// stop counting when we hit a bigger tree
						if text[nextRow][nextCol] >= text[row][column] {
							break
						}
					} else {
						// if we hit an edge
						dirDist--
						break
					}

				}
				score *= dirDist
			}

			if score > max {
				max = score
			}
		}
	}

	return max

}
