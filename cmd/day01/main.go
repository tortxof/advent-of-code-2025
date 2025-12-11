package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Rotation struct {
	Direction rune
	Distance  int
}

func ReadRotations(filename string) ([]Rotation, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var rotations []Rotation

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		direction := rune(line[0])
		if direction != 'L' && direction != 'R' {
			return nil, fmt.Errorf("invalid direction %q, expected L or R", direction)
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("failed to read int from %q: %w", line, err)
		}

		rotations = append(
			rotations,
			Rotation{Direction: direction, Distance: distance},
		)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

func main() {
	rotations, err := ReadRotations("./inputs/day01.txt")
	if err != nil {
		panic(err)
	}

	dial := 50
	zero_count := 0

	for _, rotation := range rotations {
		switch rotation.Direction {
		case 'L':
			dial -= rotation.Distance
		case 'R':
			dial += rotation.Distance
		}

		dial %= 100

		if dial == 0 {
			zero_count += 1
		}
	}

	fmt.Printf("%d\n", zero_count)
}
