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
	file, err := os.Open("./input/day10.in")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := utils.ScannerFromStart(file)

	fmt.Println(partOne(scanner))

	scanner = utils.ScannerFromStart(file)

	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) int {
	cpu := New()

	for scanner.Scan() {
		cpu.Exec(scanner.Text())

	}

	return cpu.score
}

func partTwo(scanner *bufio.Scanner) {

	cpu := New()

	for scanner.Scan() {
		cpu.Exec(scanner.Text())
	}

	cpu.Print()
}

type Cpu struct {
	clock   int
	x       int
	scoreAt int
	score   int
	crt     [6][40]string
}

func New() *Cpu {
	var crt [6][40]string
	return &Cpu{
		clock:   0,
		x:       1,
		scoreAt: 20,
		score:   0,
		crt:     crt,
	}
}

func (cpu *Cpu) calcScore() {
	if cpu.clock == cpu.scoreAt {
		cpu.score += cpu.x * cpu.clock
		cpu.scoreAt += 40
	}
}

func (cpu *Cpu) draw() {
	if cpu.x-cpu.clock%40 >= -1 && cpu.x-cpu.clock%40 <= 1 {
		// draw
		cpu.crt[cpu.clock/40][cpu.clock%40] = "#"
	} else {
		cpu.crt[cpu.clock/40][cpu.clock%40] = " "
	}
}

func (cpu *Cpu) Exec(op string) {
	opcode := op[:4]

	switch opcode {
	case "noop":
		cpu.draw()
		cpu.clock++
		cpu.calcScore()
		return
	case "addx":
		cpu.draw()
		cpu.clock++
		cpu.calcScore()
		cpu.draw()
		cpu.clock++
		cpu.calcScore()
		arg, _ := strconv.Atoi(op[5:])
		cpu.x += arg
		return
	default:
		os.Exit(1)
	}
}

func (cpu *Cpu) Print() {
	for _, v := range cpu.crt {
		fmt.Println(strings.Join(v[:], ""))
	}
}
