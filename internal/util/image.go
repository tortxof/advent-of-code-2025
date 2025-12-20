package util

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image [][]byte

// Set color of pixels in a line from pointA to pointB.
func DrawLine(img *Image, pointA, pointB Point2D, color byte) {
	maxX := max(pointA.X, pointB.X)
	minX := min(pointA.X, pointB.X)
	maxY := max(pointA.Y, pointB.Y)
	minY := min(pointA.Y, pointB.Y)
	dx := pointB.X - pointA.X
	dy := pointB.Y - pointA.Y

	if dx == 0 && dy == 0 {
		(*img)[pointB.Y][pointB.X] = color
		return
	}

	if Abs(dx) >= Abs(dy) {
		for x := minX; x <= maxX; x++ {
			y := pointA.Y + (x-pointA.X)*dy/dx
			(*img)[y][x] = color
		}
	} else {
		for y := minY; y <= maxY; y++ {
			x := pointA.X + (y-pointA.Y)*dx/dy
			(*img)[y][x] = color
		}
	}
}

func FloodFill(img *Image, point Point2D) [][]bool {
	height := len(*img)
	width := len((*img)[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	queue := []Point2D{}
	queue = append(queue, point)
	visited[point.Y][point.X] = true
	target := (*img)[point.Y][point.X]

	directions := []Point2D{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			nx := current.X + dir.X
			ny := current.Y + dir.Y

			if nx < 0 || nx >= width || ny < 0 || ny >= height {
				continue
			}

			if visited[ny][nx] {
				continue
			}

			if (*img)[ny][nx] != target {
				continue
			}

			visited[ny][nx] = true
			queue = append(queue, Point2D{X: nx, Y: ny})
		}
	}

	return visited
}

// Determine if every point in rectangle (pointA, pointB) has color in image.
func RectInArea(img *Image, pointA, pointB Point2D, color byte) bool {
	maxX := max(pointA.X, pointB.X)
	minX := min(pointA.X, pointB.X)
	maxY := max(pointA.Y, pointB.Y)
	minY := min(pointA.Y, pointB.Y)

	for _, x := range []int{minX, maxX} {
		for _, y := range []int{minY, maxY} {
			if (*img)[y][x] != color {
				return false
			}
		}
	}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if (*img)[y][x] != color {
				return false
			}
		}
	}

	return true
}

func DrawRectangle(img *Image, pointA, pointB Point2D, color byte) {
	maxX := max(pointA.X, pointB.X)
	minX := min(pointA.X, pointB.X)
	maxY := max(pointA.Y, pointB.Y)
	minY := min(pointA.Y, pointB.Y)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			(*img)[y][x] = color
		}
	}
}

func SavePng(img Image, path string) error {
	height := len(img)
	if height == 0 {
		return fmt.Errorf("image has 0 height")
	}
	width := len(img[0])
	if width == 0 {
		return fmt.Errorf("image has 0 width")
	}

	imgOut := image.NewGray(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			imgOut.Set(x, y, color.Gray{img[y][x]})
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, imgOut)
}
