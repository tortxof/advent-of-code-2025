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

func ReadData(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := ReadData("./inputs/day06.txt")
	if err != nil {
		panic(err)
	}

	col := len(data[0])

	var problems []MathProblem

	for {
		var problem MathProblem

		for {
			col--
			if col < 0 {
				break
			}
			var colData []rune
			for _, row := range data {
				colData = append(colData, row[col])
			}
			dataLen := len(colData)
			lastRune := colData[dataLen-1]
			if lastRune == '*' || lastRune == '+' {
				problem.Operator = lastRune
				colData = colData[:dataLen-1]
			}
			colStr := strings.TrimSpace(string(colData))
			if colStr == "" {
				break
			}
			value, err := strconv.Atoi(colStr)
			if err != nil {
				panic(err)
			}
			problem.Numbers = append(problem.Numbers, value)
		}
		problems = append(problems, problem)

		if col < 0 {
			break
		}
	}

	answerSum := 0
	for _, problem := range problems {
		answer, err := problem.Calculate()
		if err != nil {
			panic(err)
		}
		answerSum += answer
	}

	fmt.Println(answerSum)

}
