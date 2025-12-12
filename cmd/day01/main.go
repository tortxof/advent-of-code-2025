package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	dialSize            = 100
	initialDialPosition = 50
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

func WrapDial(dial int) int {
	if dial >= 0 && dial < dialSize {
		return dial
	}
	dial %= dialSize
	if dial < 0 {
		dial += dialSize
	}
	return dial
}

// Move `dial` in `direction` `distance` clicks, tracking how many times the
// dial crosses zero.
// Returns the new dial value and the number of zero crossings.
func CountZeroCrosses(dial int, direction rune, distance int) (int, int) {
	zeroCrossCount := 0

	for range distance {
		switch direction {
		case 'L':
			dial--
		case 'R':
			dial++
		}
		dial = WrapDial(dial)
		if dial == 0 {
			zeroCrossCount++
		}
	}

	return dial, zeroCrossCount
}

func main() {
	rotations, err := ReadRotations("./inputs/day01.txt")
	if err != nil {
		panic(err)
	}

	dial := initialDialPosition
	zeroStopCount := 0
	zeroCrossCount := 0

	var zeroCrossDuringRotation int

	for _, rotation := range rotations {
		dial, zeroCrossDuringRotation = CountZeroCrosses(
			dial,
			rotation.Direction,
			rotation.Distance,
		)

		zeroCrossCount += zeroCrossDuringRotation

		if dial == 0 {
			zeroStopCount += 1
		}
	}

	fmt.Printf("%d %d\n", zeroStopCount, zeroCrossCount)
}
