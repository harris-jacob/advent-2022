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
	file, err := os.Open("./input/day15.in")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := utils.ScannerFromStart(file)

	fmt.Println(partOne(scanner))

	scanner = utils.ScannerFromStart(file)

	fmt.Println(partTwo(scanner))
}

// construct a map of sensors and their manhattan distance
// to the closest beacon
func readInput(scanner *bufio.Scanner) map[[2]int]int {
	sensors := make(map[[2]int]int)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		sensorString, beaconString := split[0], split[1]

		sensor := parseSensor(sensorString)
		beacon := parseBeacon(beaconString)

		sensors[sensor] = manhattanDistance(sensor, beacon)

	}

	return sensors
}

func manhattanDistance(a, b [2]int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func parseSensor(sensorString string) [2]int {
	split := strings.Split(sensorString, ",")

	x, _ := strconv.Atoi(split[0][12:])
	y, _ := strconv.Atoi(split[1][3:])

	return [2]int{x, y}
}

func parseBeacon(beaconString string) [2]int {
	split := strings.Split(beaconString, ",")

	x, _ := strconv.Atoi(strings.Trim(split[0], " ")[23:])
	y, _ := strconv.Atoi(strings.Trim(split[1], " ")[2:])

	return [2]int{x, y}
}

func partOne(scanner *bufio.Scanner) int {

	sensors := readInput(scanner)
	targetY := 2000000
	impossible := make(map[int]bool)

	for sensor, distance := range sensors {
		// for each sensor there exists an interval it is
		// impossible for a beacon to exist
		Xmax := distance - abs(sensor[1]-targetY)

		interval := [2]int{sensor[0] - Xmax, sensor[0] + Xmax}

		for i := interval[0]; i < interval[1]; i++ {
			impossible[i] = true
		}
	}

	return len(impossible)
}

func partTwo(scanner *bufio.Scanner) int {
	sensors := readInput(scanner)
	directions := [][]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

	for sensor, distance := range sensors {
        for dx := 0; dx <= distance+1; dx++ {
            dy := distance -dx + 1
		for _, direction := range directions {
			pos := [2]int{sensor[0] + dx*direction[0], sensor[1] + dy*direction[1]}

			if pos[0] > 4_000_000 || pos[0] < 0 || pos[1] > 40_000_000 || pos[1] < 0{
				continue
			}

			if isValid(pos, sensors) {
				return 4_000_000*pos[0] + pos[1]
			}
		}
    }
	}

	return -1
}

func isValid(pos [2]int, sensors map[[2]int]int) bool {
	for sensor, distance := range sensors {
		if manhattanDistance(pos, sensor) <= distance {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
