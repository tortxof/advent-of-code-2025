package main

import (
	"advent-of-code-2025/internal/util"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadData(path string) ([]util.Point3D, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []util.Point3D

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		rawCoords := strings.Split(line, ",")

		if len(rawCoords) != 3 {
			return nil, fmt.Errorf("line does not contain 3 values: %q", line)
		}

		var parsedCoords []float64
		for _, coord := range rawCoords {
			val, err := strconv.ParseFloat(coord, 64)
			if err != nil {
				return nil, err
			}
			parsedCoords = append(parsedCoords, val)
		}

		points = append(
			points,
			util.Point3D{
				X: parsedCoords[0],
				Y: parsedCoords[1],
				Z: parsedCoords[2],
			},
		)

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

// Distance between a pair of points. `a` and `b` are indexes of the points.
type PointPair struct {
	distance float64
	a        int
	b        int
}

func main() {
	points, err := ReadData("./inputs/day08.txt")
	if err != nil {
		panic(err)
	}

	var distances []PointPair

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			distance := util.Distance3D(points[i], points[j])
			distances = append(distances, PointPair{distance: distance, a: i, b: j})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	uf := util.NewUnionFind(len(points))

	for _, pair := range distances[:1000] {
		uf.Union(pair.a, pair.b)
	}

	var sizes []int
	sizesSeen := util.NewSet()

	for i := range points {
		rootI := uf.Find(i)
		if !sizesSeen.Contains(rootI) {
			sizesSeen.Add(rootI)
			size := uf.GetSize(rootI)
			sizes = append(sizes, size)
		}
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	// Part 1
	fmt.Println(sizes[0] * sizes[1] * sizes[2])

	var lastPair PointPair
	for _, pair := range distances[1000:] {
		uf.Union(pair.a, pair.b)
		// if we are down to a single tree, then break
		if uf.NumSets() == 1 {
			lastPair = pair
			break
		}
	}

	// Part 2
	fmt.Println(int(points[lastPair.a].X) * int(points[lastPair.b].X))
}
