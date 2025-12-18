package main

import (
	"advent-of-code-2025/internal/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData(path string) ([]util.Point2D, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points []util.Point2D

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			return nil, fmt.Errorf("line does not contain two parts: %q", line)
		}

		var parsedParts []int
		for _, part := range parts {
			parsed, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			parsedParts = append(parsedParts, parsed)
		}

		points = append(points, util.Point2D{X: parsedParts[0], Y: parsedParts[1]})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

func ComputeArea(a, b util.Point2D) int {
	dx := a.X - b.X
	if dx < 0 {
		dx = -dx
	}
	dy := a.Y - b.Y
	if dy < 0 {
		dy = -dy
	}

	return (dx + 1) * (dy + 1)
}

func main() {
	points, err := ReadData("./inputs/day09.txt")
	if err != nil {
		panic(err)
	}

	largestRect := [3]int{0, 0, 0}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := ComputeArea(points[i], points[j])
			if area > largestRect[0] {
				largestRect = [3]int{area, i, j}
			}
		}
	}

	// Part 1
	fmt.Println(largestRect[0])
}
