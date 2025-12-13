package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func ReadInput(path string) ([]Range, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var ranges []Range

	scanner := bufio.NewScanner(file)

	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[:i], nil
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		parts := strings.Split(token, "-")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format %q: expected 2 parts, got %d", token, len(parts))
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse int %q", parts[0])
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse int %q", parts[1])
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ranges, nil
}

// Determine if a value is a valid id. The value is not valid if when split in half, the two halves are the same.
func IsValid(value int) bool {
	s := strconv.Itoa(value)
	if len(s)%2 != 0 {
		return true
	}
	first := s[:len(s)/2]
	last := s[len(s)/2:]
	return !(first == last)
}

// Split a string into parts of length `l`. Returns a slice of strings.
func SplitParts(input string, l int) []string {
	var parts []string
	for offset := 0; offset < len(input); offset += l {
		parts = append(parts, input[offset:offset+l])
	}
	return parts
}

// Returns `true` if `input` is made up only of repeated `pattern`.
func PartsEqual(input string, pattern string) bool {
	patternLen := len(pattern)
	for offset := patternLen; offset < len(input); offset += patternLen {
		if input[offset:offset+patternLen] != pattern {
			return false
		}
	}
	return true
}

// `value` is not valid only if it is made entirely from a repeated pattern.
func IsValid2(value int) bool {
	s := strconv.Itoa(value)
	length := len(s)

	for partLen := 1; partLen <= length/2; partLen++ {
		if length%partLen == 0 {
			if PartsEqual(s, s[:partLen]) {
				return false
			}
		}
	}
	return true
}

func main() {
	ranges, err := ReadInput("./inputs/day02.txt")
	if err != nil {
		panic(err)
	}

	invalidSum := 0
	for _, value := range ranges {
		for i := value.Start; i <= value.End; i++ {
			if !IsValid(i) {
				invalidSum += i
			}
		}
	}

	fmt.Println(invalidSum)

	invalidSum2 := 0
	for _, value := range ranges {
		for i := value.Start; i <= value.End; i++ {
			if !IsValid2(i) {
				invalidSum2 += i
			}
		}
	}

	fmt.Println(invalidSum2)
}
