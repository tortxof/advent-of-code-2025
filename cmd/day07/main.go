package main

import (
	"bufio"
	"fmt"
	"os"
)

var validTachyonRunes = map[rune]bool{
	'S': true,
	'.': true,
	'^': true,
}

type TachyonManifold struct {
	Levels []TachyonLevel
}

type TachyonLevel struct {
	Cells []rune
}

type BeamSet struct {
	Positions map[int]struct{}
}

func NewBeamSet() *BeamSet {
	return &BeamSet{
		Positions: make(map[int]struct{}),
	}
}

func (b *BeamSet) Add(position int) {
	b.Positions[position] = struct{}{}
}

func (b *BeamSet) Remove(position int) {
	delete(b.Positions, position)
}

func (b *BeamSet) Contains(position int) bool {
	_, exists := b.Positions[position]
	return exists
}

func (b *BeamSet) Len() int {
	return len(b.Positions)
}

func EvaluateTachyonLevel(beamSet *BeamSet, level TachyonLevel) int {
	numSplits := 0
	for i, r := range level.Cells {
		switch r {
		case 'S':
			beamSet.Add(i)
		case '^':
			if beamSet.Contains(i) {
				numSplits++
				beamSet.Remove(i)
				beamSet.Add(i - 1)
				beamSet.Add(i + 1)
			}
		}
	}
	return numSplits
}

func ReadData(path string) (TachyonManifold, error) {
	file, err := os.Open(path)
	if err != nil {
		return TachyonManifold{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	manifold := TachyonManifold{}

	for scanner.Scan() {
		line := scanner.Text()
		level := TachyonLevel{}
		for _, r := range line {
			if !validTachyonRunes[r] {
				return manifold, fmt.Errorf("unknown symbol in input: %v", r)
			}
			level.Cells = append(level.Cells, r)
		}
		manifold.Levels = append(manifold.Levels, level)
	}

	if err := scanner.Err(); err != nil {
		return manifold, err
	}

	return manifold, nil
}

func main() {
	manifold, err := ReadData("./inputs/day07.txt")
	if err != nil {
		panic(err)
	}

	beamSet := NewBeamSet()
	numSplits := 0
	for _, level := range manifold.Levels {
		var levelSplits int
		levelSplits = EvaluateTachyonLevel(beamSet, level)
		numSplits += levelSplits
	}

	// Part 1: 1518
	fmt.Println(numSplits)
}
