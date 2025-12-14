package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Bank []int

func (bank *Bank) Add(value int) {
	*bank = append(*bank, value)
}

func (bank Bank) ConcatDigits() int {
	sum := 0
	for _, digit := range bank {
		sum = sum*10 + digit
	}
	return sum
}

// Bank with `x` digits trimmed from right.
func (bank Bank) TrimRight(x int) Bank {
	return bank[:len(bank)-x]
}

func (bank Bank) MaxJoltage(numBatteries int) int {
	remainingBank := make(Bank, len(bank))
	copy(remainingBank, bank)
	var digits Bank
	for battIndex := range numBatteries {
		searchBank := remainingBank.TrimRight(numBatteries - 1 - battIndex)
		i, digit := searchBank.FindLargestDigit()
		digits = append(digits, digit)
		remainingBank = remainingBank[i+1:]
	}
	return digits.ConcatDigits()
}

// Find the largest digit in a bank. Returns the index of the digit and the digit.
func (bank Bank) FindLargestDigit() (int, int) {
	for searchDigit := 9; searchDigit >= 0; searchDigit-- {
		for i := range bank {
			if bank[i] == searchDigit {
				return i, searchDigit
			}
		}
	}
	panic("All elements in bank were negative.")
}

func ReadData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var banks []Bank

	for scanner.Scan() {
		line := scanner.Text()
		var bank Bank
		for _, joltageRune := range line {
			joltage, err := strconv.Atoi(string(joltageRune))
			if err != nil {
				return nil, err
			}
			bank.Add(joltage)
		}
		banks = append(banks, bank)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, nil
}

func main() {
	banks, err := ReadData("./inputs/day03.txt")
	if err != nil {
		panic(err)
	}

	totalJoltage := 0
	for _, bank := range banks {
		// Part 1. 2 batteries.
		totalJoltage += bank.MaxJoltage(2)
	}

	fmt.Println(totalJoltage)

	totalJoltage2 := 0
	for _, bank := range banks {
		// Part 2. 12 batteries.
		totalJoltage2 += bank.MaxJoltage(12)
	}

	fmt.Println(totalJoltage2)
}
