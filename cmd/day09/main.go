package main

import (
	"advent-of-code-2025/internal/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData(path string) ([]util.Point2D, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points []util.Point2D

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			return nil, fmt.Errorf("line does not contain two parts: %q", line)
		}

		var parsedParts []int
		for _, part := range parts {
			parsed, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			parsedParts = append(parsedParts, parsed)
		}

		points = append(points, util.Point2D{X: parsedParts[0], Y: parsedParts[1]})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

func ComputeArea(a, b util.Point2D) int {
	dx := a.X - b.X
	if dx < 0 {
		dx = -dx
	}
	dy := a.Y - b.Y
	if dy < 0 {
		dy = -dy
	}

	return (dx + 1) * (dy + 1)
}

func main() {
	points, err := ReadData("./inputs/day09.txt")
	if err != nil {
		panic(err)
	}

	largestRect := [3]int{0, 0, 0}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := ComputeArea(points[i], points[j])
			if area > largestRect[0] {
				largestRect = [3]int{area, i, j}
			}
		}
	}

	// Part 1
	fmt.Println(largestRect[0])

	fmt.Println("Calculating max size...")
	var maxX int
	var maxY int
	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}
	maxX += 2
	maxY += 2

	fmt.Println("Initializing Image...")
	floorTiles := make(util.Image, maxY)
	for i := range floorTiles {
		floorTiles[i] = make([]byte, maxX)
	}

	fmt.Println("Drawing lines...")
	var lastPoint util.Point2D
	firstIteration := true
	for _, point := range points {
		if firstIteration {
			firstIteration = false
			lastPoint = point
			continue
		}
		util.DrawLine(&floorTiles, lastPoint, point, 255)
		lastPoint = point
	}
	util.DrawLine(&floorTiles, lastPoint, points[0], 255)

	fmt.Println("Flood fill select...")
	outside := util.FloodFill(&floorTiles, util.Point2D{X: 0, Y: 0})

	fmt.Println("Flood fill apply...")
	for y := range outside {
		for x := range outside[y] {
			if !outside[y][x] {
				floorTiles[y][x] = 255
			}
		}
	}

	fmt.Println("Looking for largest rectangle...")
	largestRectInside := [3]int{0, 0, 0}
	numPoints := len(points)
	totalPairs := numPoints * (numPoints - 1) / 2
	pairsChecked := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			pairsChecked++
			area := ComputeArea(points[i], points[j])
			if area > largestRectInside[0] {
				if util.RectInArea(&floorTiles, points[i], points[j], 255) {
					largestRectInside = [3]int{area, i, j}
				}
			}
		}
		progressPercent := (float32(pairsChecked) / float32(totalPairs)) * 100
		fmt.Printf("\rProgress %.2f%% ", progressPercent)
	}
	fmt.Println()

	// Part 2
	fmt.Println(largestRectInside[0])

	fmt.Println("Drawing rectangle...")
	util.DrawRectangle(
		&floorTiles,
		points[largestRectInside[1]],
		points[largestRectInside[2]],
		127,
	)

	fmt.Println("Saving PNG...")
	util.SavePng(floorTiles, "day09.png")
}
