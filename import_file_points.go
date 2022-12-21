package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

const LinesCount int = 18

func import_file_points_by_path(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return import_file_points_stream(f)
}

func import_file_points_stream(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0, LinesCount)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return lines
}
