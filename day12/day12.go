package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("./input/day12.in")

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
    hill := make([][]rune, 0)
    queue := make([][3]int, 0)
    target := [2]int{}

    // append our starting point to queue
    // find our target square
    x := 0
    for scanner.Scan() {
        row := make([]rune, 0)
        for y, v := range scanner.Text() {
            row = append(row, v)
            if v == 'S' {
                queue = append(queue, [3]int{x,y, 0})
            } else if v == 'E' {
                target = [2]int{x,y}
            }
        }

        hill = append(hill, row)
        x ++
    }


    return dfs(hill, queue, target)
}


func partTwo(scanner *bufio.Scanner) int {
    hill := make([][]rune, 0)
    queue := make([][3]int, 0)
    target := [2]int{}

    // append all starting squares to dfs queue
    // find the target square
    x := 0
    for scanner.Scan() {
        row := make([]rune, 0)
        for y, v := range scanner.Text() {
            row = append(row, v)
            if v == 'S' || v == 'a' {
                queue = append(queue, [3]int{x,y, 0})
            } else if v == 'E' {
                target = [2]int{x,y}
            }
        }

        hill = append(hill, row)
        x ++
    }


    return dfs(hill, queue, target)
}



func dfs(hill [][]rune, queue [][3]int, target [2]int) int {
    visited := make(map[[2]int]bool)
    
    for len(queue) != 0 {
        X := queue[0][0]
        Y := queue[0][1]
        count := queue[0][2]
        queue = queue[1:]

        if target[0] == X && target[1] == Y {
            return count
        }

        // visit neighbours
        for _, v := range []string{"N", "E", "S", "W"} {
            dir := evalDirection(v)
            newX := X + dir[0]
            newY := Y + dir[1]


            if newX < len(hill) && newY < len(hill[0]) && newX >= 0 && newY >= 0 {
                ok := visited[[2]int{newX, newY}]
                if evalLetter(hill[newX][newY]) - evalLetter(hill[X][Y]) <= 1 && !ok {
                    coord := [3]int{newX, newY, count+1}
                    queue = append(queue, coord)
                    visited[[2]int{coord[0], coord[1]}] = true
                }
            }
        }

    }

	return -1
}


func evalDirection(dir string) [2]int {
    switch dir {
    case "N":
        return [2]int{1, 0}
    case "S" :
        return [2]int{-1, 0}
    case "E" :
        return [2]int{0, 1}
    default:
        return [2]int{0, -1}

    }
}

func evalLetter(letter rune) int {
    switch rune(letter) {
    case 'E':
        return int('z')

    case 'S':
        return int('a')

    default:
        return int(letter)
    }
}

