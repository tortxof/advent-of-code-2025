package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func ReadData(path string) ([]Machine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var machines []Machine

	for scanner.Scan() {
		line := scanner.Text()

		machines = append(machines, ParseMachine(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

func main() {
	machines, err := ReadData("./inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	var totalButtonsPushed = 0
	for _, machine := range machines {
		numButtons := len(machine.Buttons)
		leastNumButtonsPushed := numButtons + 1
		for buttonMask := range 1 << numButtons {
			numButtonsPushed := bits.OnesCount(uint(buttonMask))
			if numButtonsPushed >= leastNumButtonsPushed {
				continue
			}
			// use buttonMask to get a sub-selection of buttons from machine.Buttons
			var selectedButtons []uint
			for i, button := range machine.Buttons {
				if buttonMask&(1<<i) != 0 {
					selectedButtons = append(selectedButtons, button)
				}
			}
			lightState := CalcLightState(selectedButtons)
			if lightState == machine.DesiredLightState {
				leastNumButtonsPushed = numButtonsPushed
			}
		}
		totalButtonsPushed += leastNumButtonsPushed
	}

	fmt.Println(totalButtonsPushed)
}
