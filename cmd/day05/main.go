package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func NewRange(a, b int) Range {
	return Range{Start: min(a, b), End: max(a, b)}
}

// Check if value is in the range. Ranges are inclusive.
func (r Range) InRange(value int) bool {
	return value >= r.Start && value <= r.End
}

func GetRangeBounds(ranges []Range) (minStart, maxEnd int) {
	if len(ranges) == 0 {
		return 0, 0
	}

	minStart = ranges[0].Start
	maxEnd = ranges[0].End

	for _, r := range ranges {
		if r.Start < minStart {
			minStart = r.Start
		}
		if r.End > maxEnd {
			maxEnd = r.End
		}
	}

	return minStart, maxEnd
}

func ReadData(path string) (ranges []Range, availableIds []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			err = fmt.Errorf(
				"invalid range format %q: expected 2 parts, got %d",
				line,
				len(parts),
			)
			return
		}

		var a, b int
		a, err = strconv.Atoi(parts[0])
		if err != nil {
			return
		}
		b, err = strconv.Atoi(parts[1])
		if err != nil {
			return
		}

		ranges = append(ranges, NewRange(a, b))
	}

	for scanner.Scan() {
		line := scanner.Text()

		var id int
		id, err = strconv.Atoi(line)
		if err != nil {
			return
		}

		availableIds = append(availableIds, id)
	}

	err = scanner.Err()

	return
}

func main() {
	ranges, availableIds, err := ReadData("./inputs/day05.txt")
	if err != nil {
		panic(err)
	}

	minStart, maxEnd := GetRangeBounds(ranges)

	numFresh := 0

	// Check each of the available ingredient IDs.
	for _, id := range availableIds {
		// End iteration early if ID is outside the super-range covered by all ranges.
		if id < minStart || id > maxEnd {
			continue
		}
		// Check each range to see if ID is inside. Increment numFresh and break
		// early if we find a covering range.
		for _, r := range ranges {
			if r.InRange(id) {
				numFresh++
				break
			}
		}
	}

	// Part 1
	fmt.Println(numFresh)

	// Sort ranges by `Start`.
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.Start - b.Start
	})

	highestId := 0
	numFresh = 0
	for _, r := range ranges {
		if r.End <= highestId {
			continue
		}
		start := max(highestId+1, r.Start)
		numFresh += r.End - start + 1
		highestId = r.End
	}

	// Part 2
	fmt.Println(numFresh)
}
