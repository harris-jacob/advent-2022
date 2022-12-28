package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day13.in")
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
	packets := [][2][]interface{}{}

	// read input
	for scanner.Scan() {
		lhs := scanner.Text()
		scanner.Scan()
		rhs := scanner.Text()
		packets = append(packets, parsePacketPair(lhs, rhs))
		scanner.Scan()
	}

	// compare packets
	sum := 0
	for idx, packet := range packets {
		result := compare(packet[0], packet[1])

		if result == RightOrder {
			sum += idx + 1
		}
	}

	return sum
}

func partTwo(scanner *bufio.Scanner) int {
	packets := [][]interface{}{}

	// read input
	for scanner.Scan() {
		packet := scanner.Text()
		packets = append(packets, parsePacket(packet))
		scanner.Scan()
		packet = scanner.Text()
		packets = append(packets, parsePacket(packet))
		scanner.Scan()
	}

	// put in divider packets
	divider1 := []interface{}{[]interface{}{6.0}}
	divider2 := []interface{}{[]interface{}{2.0}}

	packets = append(packets, divider1)
	packets = append(packets, divider2)

	// sort input
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == RightOrder
	})


	idx1 := 0
	idx2 := 0
	// find new position of divider packets
	for idx, packet := range packets {
		if len(packet) == 1 {
			packetArr, isArray := packet[0].([]interface{})
			if isArray && len(packetArr) == 1 {
				asInt, isInt := packetArr[0].(float64)

				if isInt && int(asInt) == 6 {
					idx1 = idx + 1
				} else if isInt && int(asInt) == 2 {
					idx2 = idx + 1
				}
			}
		}

	}

	return idx1 * idx2
}

func parsePacketPair(lhs, rhs string) [2][]interface{} {
	lhsList := parsePacket(lhs)
	rhsList := parsePacket(rhs)

	return [2][]interface{}{lhsList, rhsList}

}

func parsePacket(packet string) []interface{} {
	parsed := []interface{}{}
	json.Unmarshal([]byte(packet), &parsed)

	return parsed
}

func compare(lhs, rhs interface{}) ComparisonResult {
	// convert to integers
	lhsInt, isLhsInt := lhs.(float64)
	rhsInt, isRhsInt := rhs.(float64)

	if isLhsInt && isRhsInt {
		if lhsInt < rhsInt {
			return RightOrder
		} else if int(lhsInt) == int(rhsInt) {
			return CompareNext
		} else {
			return WrongOrder
		}
	}

	// else one of the values is a list
	lhsList, isLhsList := lhs.([]interface{})
	rhsList, isRhsList := rhs.([]interface{})

	if isLhsList && isRhsList {
		for i := 0; i < len(lhsList) && i < len(rhsList); i++ {
			result := compare(lhsList[i], rhsList[i])

			if result == RightOrder || result == WrongOrder {
				return result
			}
		}
		// special rules for the case of unequal length lists
		if len(lhsList) < len(rhsList) {
			return RightOrder
		} else if len(rhsList) < len(lhsList) {
			return WrongOrder
		} else {
			return CompareNext
		}
	}

	// one is a list and the other is an integers
	// so we convert the integer to a list of length 1
	// and compare the lists
	if isLhsList {
		return compare(lhsList, []interface{}{rhsInt})
	}

	return compare([]interface{}{lhsInt}, rhsList)

}

type ComparisonResult int

const (
	RightOrder ComparisonResult = iota
	CompareNext
	WrongOrder
)
