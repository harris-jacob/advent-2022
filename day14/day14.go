package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day14.in")

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
    rocks := make(map[[2]int]bool, 0)
    for scanner.Scan() {
        parseLine(scanner.Text(), rocks)
    }

    floor := maxY(rocks)
    sand := [2]int{500, 0}
    count := 0
    for i := 0; i< 10000000; i++ {
        if sand[1] >= floor {
            break
        }

        // can sand move down
        if _, ok :=rocks[[2]int{sand[0], sand[1]+1}]; !ok {
            sand[1]++
        // tries to move diagnonally left
        } else if _, ok :=rocks[[2]int{sand[0]-1, sand[1]+1}]; !ok {
                sand[0]--
                sand[1]++
        // tries to move diagonally right
        } else if _, ok := rocks[[2]int{sand[0]+1, sand[1]+1}]; !ok {
            sand[0]++
            sand[1]++
        // stuck so becomes rock
        } else {
            count++
            rocks[sand] = true
            sand = [2]int{500, 0}
        }
    }

	return count
}

func partTwo(scanner *bufio.Scanner) int {
    rocks := make(map[[2]int]bool, 0)
    for scanner.Scan() {
        parseLine(scanner.Text(), rocks)
    }
    floor := maxY(rocks) + 2

    sand := [2]int{500, 0}
    count := 0
    for i := 0; i< 10000000; i++ {
        if _, ok := rocks[[2]int{500, 0}]; ok {
            break
        }

        // now if we hit the floor, we add the rock
        if sand[1]+1 >= floor {
            count++
            rocks[sand] = true
            sand = [2]int{500, 0}
            continue
        }

        // can sand move down
        if _, ok :=rocks[[2]int{sand[0], sand[1]+1}]; !ok {
            sand[1]++
        // tries to move diagnonally left
        } else if _, ok :=rocks[[2]int{sand[0]-1, sand[1]+1}]; !ok {
                sand[0]--
                sand[1]++
        // tries to move diagonally right
        } else if _, ok := rocks[[2]int{sand[0]+1, sand[1]+1}]; !ok {
            sand[0]++
            sand[1]++
        // stuck so becomes rock
        } else {
            count++
            rocks[sand] = true
            sand = [2]int{500, 0}
        }
    }

	return count
}


func parseLine(line string, rocks map[[2]int]bool) {
    var coords [][]int
    for _, v:= range strings.Split(line, "->") {
        coord := strings.Split(strings.Trim(v, " "), ",")
        x, _ := strconv.Atoi(coord[0])
        y, _ := strconv.Atoi(coord[1])
        coords = append(coords, []int{x, y})
    }

    j:=1
    for i:=0; j<len(coords); i++ {
        x1, y1 := coords[i][0], coords[i][1]
        x2, y2 := coords[j][0], coords[j][1]

        dx := x2-x1
        dy := y2-y1

        delta := math.Max(math.Abs(float64(dx)), math.Abs(float64(dy)))

        for i:=0;i<=int(delta); i++ {
            if(dx > 0) {
                rocks[[2]int{x1+i, y1}] = true
            } else if (dx < 0) {
                rocks[[2]int{x1-i, y1}] = true
            } else if (dy > 0) {
                rocks[[2]int{x1, y1+i}] = true
            } else if (dy < 0) {
                rocks[[2]int{x1, y1-i}] = true
            }
        }

        j++
    }
}


// find max y value in rocks
func maxY(rocks map[[2]int]bool) int {
    max := 0
    for k := range rocks {
        if k[1] > max {
            max = k[1]
        }
    }
    return max
}
