package main

import (
	"bufio"
	"io"
)

const LinesCount int = 18

func import_file_points_stream(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0, LinesCount)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return lines
}
