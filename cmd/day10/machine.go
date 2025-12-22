package main

import "advent-of-code-2025/internal/util"

func CalcLightState(buttons []uint) (lights uint) {
	for _, button := range buttons {
		lights = lights ^ button
	}
	return lights
}

type Machine struct {
	DesiredLightState uint
	Buttons           []uint
	Joltages          []int
}

type MachineParser struct {
	state    func(*MachineParser, rune)
	lights   uint
	buttons  []uint
	joltages []int

	numLights          int
	currentNumber      int
	currentNumberValid bool
	currentButton      uint
}

func NewMachineParser() *MachineParser {
	machineParser := &MachineParser{state: ProcessRuneRoot}

	return machineParser
}

func ParseMachine(line string) Machine {
	parser := NewMachineParser()
	for _, r := range line {
		parser.state(parser, r)
	}
	return Machine{
		DesiredLightState: util.ReverseBits(parser.lights, parser.numLights),
		Buttons:           parser.buttons,
		Joltages:          parser.joltages,
	}
}

func ProcessRuneRoot(p *MachineParser, r rune) {
	switch r {
	case '[':
		p.numLights = 0
		p.state = ProcessRuneLights
	case '(':
		p.currentNumber = 0
		p.currentNumberValid = false
		p.currentButton = 0
		p.state = ProcessRuneButton
	case '{':
		p.currentNumber = 0
		p.currentNumberValid = false
		p.state = ProcessRuneJoltages
	}
}

func ProcessRuneLights(p *MachineParser, r rune) {
	switch r {
	case ']':
		p.state = ProcessRuneRoot
	case '.':
		p.lights <<= 1
		p.numLights++
	case '#':
		p.lights <<= 1
		p.lights++
		p.numLights++
	}
}

func ProcessRuneButton(p *MachineParser, r rune) {
	switch {
	case r == ')':
		if p.currentNumberValid {
			var buttonBit uint = 1 << p.currentNumber
			p.currentButton = p.currentButton | buttonBit
		}
		if p.currentButton != 0 {
			p.buttons = append(p.buttons, p.currentButton)
		}
		p.currentNumber = 0
		p.currentNumberValid = false
		p.currentButton = 0
		p.state = ProcessRuneRoot
	case r == ',':
		if p.currentNumberValid {
			var buttonBit uint = 1 << p.currentNumber
			p.currentButton = p.currentButton | buttonBit
		}
		p.currentNumber = 0
		p.currentNumberValid = false
	case util.IsDigit(r):
		digit := util.ParseAsciiDigit(r)
		p.currentNumber = p.currentNumber*10 + digit
		p.currentNumberValid = true
	}
}

func ProcessRuneJoltages(p *MachineParser, r rune) {
	switch {
	case r == '}':
		if p.currentNumberValid {
			p.joltages = append(p.joltages, p.currentNumber)
		}
		p.currentNumber = 0
		p.currentNumberValid = false
		p.state = ProcessRuneRoot
	case r == ',':
		if p.currentNumberValid {
			p.joltages = append(p.joltages, p.currentNumber)
		}
		p.currentNumber = 0
		p.currentNumberValid = false
	case util.IsDigit(r):
		digit := util.ParseAsciiDigit(r)
		p.currentNumber = p.currentNumber*10 + digit
		p.currentNumberValid = true
	}
}
