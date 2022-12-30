package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

const LinesCount int = 18

func ImportFilePointsByPath(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return ImportFilePointsStream(f)
}

func ImportFilePointsStream(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0, LinesCount)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return lines
}
