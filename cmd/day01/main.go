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

// Move `dial` in `direction` `distance` clicks, tracking how many times the
// dial crosses zero.
// Returns the new dial value and the number of zero crossings.
func count_zero_crosses(dial int, direction rune, distance int) (int, int) {
	zero_cross_count := 0

	for range distance {
		switch direction {
		case 'L':
			dial--
		case 'R':
			dial++
		}
		dial %= 100
		if dial == 0 {
			zero_cross_count++
		}
	}

	return dial, zero_cross_count
}

func main() {
	rotations, err := ReadRotations("./inputs/day01.txt")
	if err != nil {
		panic(err)
	}

	dial := 50
	zero_stop_count := 0
	zero_cross_count := 0

	var zero_cross_during_rotation int

	for _, rotation := range rotations {
		dial, zero_cross_during_rotation = count_zero_crosses(
			dial,
			rotation.Direction,
			rotation.Distance,
		)

		zero_cross_count += zero_cross_during_rotation

		if dial == 0 {
			zero_stop_count += 1
		}
	}

	fmt.Printf("%d %d\n", zero_stop_count, zero_cross_count)
}
