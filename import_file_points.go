package main

import (
	"bufio"
	"log"
	"os"
)

const LinesCount int = 18

func importFilePoints() []string {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0, LinesCount)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	f.Close()

	return lines
}
