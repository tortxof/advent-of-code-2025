package util

import "math"

type Point2D struct {
	X int
	Y int
}

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func Distance3D(a, b Point3D) float64 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	dz := b.Z - a.Z

	squares := dx*dx + dy*dy + dz*dz

	return math.Sqrt(squares)
}

type Set struct {
	members map[int]struct{}
}

func NewSet() *Set {
	return &Set{members: make(map[int]struct{})}
}

func (s *Set) Add(x int) {
	s.members[x] = struct{}{}
}

func (s *Set) Remove(x int) {
	delete(s.members, x)
}

func (s *Set) Contains(x int) bool {
	_, exists := s.members[x]
	return exists
}

func (s *Set) Len() int {
	return len(s.members)
}

func (s *Set) Members() []int {
	members := make([]int, 0, len(s.members))
	for k := range s.members {
		members = append(members, k)
	}
	return members
}
