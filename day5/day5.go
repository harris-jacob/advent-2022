package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/harrisja-jacob/advent-2022/utils"
)

func main() {
	file, err := os.Open("../input/day5.in")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := utils.ScannerFromStart(file)

	fmt.Println(partOne(scanner))

	scanner = utils.ScannerFromStart(file)

	fmt.Println(partTwo(scanner))
}

func partOne(scanner *bufio.Scanner) string {
	stacks := parseStacks(scanner)

	mover9000(scanner, stacks)

	return topLine(stacks)

}

func partTwo(scanner *bufio.Scanner) string {

	stacks := parseStacks(scanner)

	mover9001(scanner, stacks)

	return topLine(stacks)

}

func topLine(stacks map[int]*Stack) string {

	ans := ""
	for i := 0; i < len(stacks); i++ {
		ans += stacks[i].Pop()
	}

	return ans
}

func mover9000(scanner *bufio.Scanner, stacks map[int]*Stack) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		instruction := parseInstruction(line)

		fromStack := stacks[instruction.from]
		toStack := stacks[instruction.to]

		for i := 0; i < instruction.moves; i++ {
			val := fromStack.Pop()
			toStack.Push(val)
		}
	}
}

func mover9001(scanner *bufio.Scanner, stacks map[int]*Stack) {
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		instruction := parseInstruction(line)

		fromStack := stacks[instruction.from]
		toStack := stacks[instruction.to]

		toStack.TransferFrom(fromStack, instruction.moves)

	}

}

type Instruction struct {
	to    int
	from  int
	moves int
}

func parseInstruction(line string) *Instruction {
	split := strings.Split(line, " ")

	moves, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])

	return &Instruction{
		moves: moves,
		// account for zero
		from: from - 1,
		to:   to - 1,
	}
}

func parseStacks(scanner *bufio.Scanner) map[int]*Stack {
	stacks := make(map[int]*Stack)

	// parse stacks
	for scanner.Scan() {
		// look for line with the stacks
		line := scanner.Text()

		if strings.Contains(line, "1") {
			break
		}

		for i := 0; i < len(line); i += 4 {
			r, _ := utf8.DecodeRuneInString(line[i+1:])
			s := string(r)
			if s != " " {
				stack, ok := stacks[i/4]
				if !ok {
					stacks[i/4] = &Stack{}
					stack = stacks[i/4]
				}

				stack.Prepend(s)
			}
		}

	}

	return stacks
}

type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)

}

func (s *Stack) Prepend(val string) {
	*s = append([]string{val}, *s...)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (to *Stack) TransferFrom(from *Stack, n int) {
	index := len(*from) - n

	for i := index; i < len(*from); i++ {
		*to = append(*to, (*from)[i])
	}

	*from = (*from)[0:index]

}
