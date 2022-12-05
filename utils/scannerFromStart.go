package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ScannerFromStart(file *os.File) *bufio.Scanner {

	file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner
}
