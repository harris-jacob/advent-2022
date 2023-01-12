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
	file, err := os.Open("./input/day18.in")

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
	cubes := make(map[[3]int]bool, 0)

	for scanner.Scan() {
		cube := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(cube[0])
		y, _ := strconv.Atoi(cube[1])
		z, _ := strconv.Atoi(cube[2])

		cubes[[3]int{x, y, z}] = true
	}

	for cube := range cubes {
		for _, direction := range directions() {
			pos := [3]int{cube[0] + direction[0], cube[1] + direction[1], cube[2] + direction[2]}

			if _, ok := cubes[pos]; !ok {
				// there isnt an adjacent cube
				count++
			}
		}

	}

	return count
}

func partTwo(scanner *bufio.Scanner) int {
	count := 0
	cubes := make(map[[3]int]bool, 0)

	for scanner.Scan() {
		cube := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(cube[0])
		y, _ := strconv.Atoi(cube[1])
		z, _ := strconv.Atoi(cube[2])

		cubes[[3]int{x, y, z}] = true
	}

	for cube := range cubes {
		for _, direction := range directions() {
			pos := [3]int{cube[0] + direction[0], cube[1] + direction[1], cube[2] + direction[2]}
			if _, ok := cubes[pos]; !ok && isExposed(pos, cubes) {
				count++
			}
		}
	}

	return count
}

func isExposed(cube [3]int, cubes map[[3]int]bool) bool {
	toSearch := NewStack()
	toSearch.Push(cube)

	seen := make(map[[3]int]bool, 0)

	for toSearch.Len() > 0 {

		current := toSearch.Pop()

		if _, ok := cubes[current]; ok {
			continue
		}

		if isOutsidebounds(current) {
			return true
		}

		if _, ok := seen[current]; ok {
			continue
		}

		seen[current] = true

		for _, direction := range directions() {
			pos := [3]int{current[0] + direction[0], current[1] + direction[1], current[2] + direction[2]}

			toSearch.Push(pos)
		}

	}

	return false
}

func directions() [][3]int {
	return [][3]int{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}
}

func isOutsidebounds(pos [3]int) bool {
	if pos[0] < 0 || pos[0] >= 20 {
		return true
	}
	if pos[1] < 0 || pos[1] >= 20 {
		return true
	}
	if pos[2] < 0 || pos[2] >= 20 {
		return true
	}

	return false
}

type Stack struct {
	items [][3]int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(i [3]int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() [3]int {
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}

func (s *Stack) Len() int {
	return len(s.items)
}
