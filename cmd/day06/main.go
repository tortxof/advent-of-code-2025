package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MathProblem struct {
	Numbers  []int
	Operator rune
}

func (m MathProblem) Calculate() (answer int, err error) {
	switch m.Operator {
	case '*':
		answer = 1
	case '+':
		answer = 0
	default:
		return answer, fmt.Errorf("unknown operator: %q", m.Operator)
	}

	for _, v := range m.Numbers {
		switch m.Operator {
		case '*':
			answer *= v
		case '+':
			answer += v
		}
	}

	return answer, nil
}

func ReadData(path string) ([]MathProblem, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var mathProblems []MathProblem

	scanner := bufio.NewScanner(file)

	// Get the first element of each problem.
	if scanner.Scan() {
		line := scanner.Text()

		for token := range strings.SplitSeq(line, " ") {
			token = strings.TrimSpace(token)
			if token == "" {
				continue
			}
			value, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}
			mathProblems = append(
				mathProblems,
				MathProblem{Numbers: []int{value}, Operator: ' '},
			)
		}
	}

	// Get the rest of the elements and add to each problem.
	for scanner.Scan() {
		line := scanner.Text()

		problemIndex := 0
		for token := range strings.SplitSeq(line, " ") {
			token = strings.TrimSpace(token)
			if token == "" {
				continue
			}
			if token == "*" || token == "+" {
				mathProblems[problemIndex].Operator = []rune(token)[0]
				problemIndex++
				continue
			}
			value, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}
			mathProblems[problemIndex].Numbers = append(mathProblems[problemIndex].Numbers, value)
			problemIndex++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mathProblems, nil
}

func main() {
	mathProblems, err := ReadData("./inputs/day06.txt")
	if err != nil {
		panic(err)
	}

	answerSum := 0
	for _, problem := range mathProblems {
		answer, err := problem.Calculate()
		if err != nil {
			panic(err)
		}
		answerSum += answer
	}

	fmt.Println(answerSum)
}
