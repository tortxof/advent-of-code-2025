package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
	cells [][]bool
}

type Coord struct {
	X, Y int
}

func (g *Grid) AppendRow(row []bool) {
	g.cells = append(g.cells, row)
}

func (g *Grid) UnsetCell(coord Coord) {
	if !g.InBounds(coord) {
		return
	}
	g.cells[coord.Y][coord.X] = false
}

func (g Grid) InBounds(coord Coord) bool {
	if coord.Y < 0 || coord.Y >= len(g.cells) {
		return false
	}
	if coord.X < 0 || coord.X >= len(g.cells[coord.Y]) {
		return false
	}
	return true
}

func (g Grid) Get(coord Coord) bool {
	if !g.InBounds(coord) {
		return false
	}
	return g.cells[coord.Y][coord.X]
}

// Count number of occupied cells in an area centered on x, y. Also counts the
// center cell, x, y. The search area is always square. Cells coordinate x plus
// or minus radius and y plus or minus radius are counted.
func (g Grid) CountArea(coord Coord, radius int) int {
	occupiedCount := 0
	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			if g.Get(Coord{X: coord.X + dx, Y: coord.Y + dy}) {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

// Find all occupied cells with fewer than `threshold` occupied neighbors.
// Return number of cells found, and a slice of `Coord` of those cells.
func (g Grid) GetAvailableCells(threshold int) (int, []Coord) {
	availableCells := 0
	var coords []Coord

	for y := range g.cells {
		for x := range g.cells[y] {
			coord := Coord{X: x, Y: y}
			if !g.Get(coord) {
				continue
			}
			numNeighbors := g.CountArea(coord, 1) - 1
			if numNeighbors < threshold {
				availableCells++
				coords = append(coords, coord)
			}
		}
	}

	return availableCells, coords
}

func ParseLine(line string) []bool {
	var row []bool
	for _, cell := range line {
		row = append(row, cell == '@')
	}
	return row
}

func ReadData(path string) (Grid, error) {
	var grid Grid

	file, err := os.Open(path)
	if err != nil {
		return grid, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid.AppendRow(ParseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return grid, err
	}

	return grid, nil
}

func main() {
	grid, err := ReadData("./inputs/day04.txt")
	if err != nil {
		panic(err)
	}

	threshold := 4
	totalAvailableCells := 0
	firstIteration := true
	for {
		availableCells, coords := grid.GetAvailableCells(threshold)
		if availableCells == 0 {
			break
		}
		if firstIteration {
			firstIteration = false
			// Part 1: 1523
			fmt.Println(availableCells)
		}
		totalAvailableCells += availableCells
		for _, coord := range coords {
			grid.UnsetCell(coord)
		}
	}

	// Part 2: 9290
	fmt.Println(totalAvailableCells)
}
