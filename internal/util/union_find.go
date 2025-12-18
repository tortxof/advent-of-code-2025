package util

import "fmt"

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}

	return true
}

func (uf *UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) Inspect() {
	fmt.Printf("parent: %v\n", uf.parent)
	fmt.Printf("size:   %v\n", uf.size)
}

// Count number of distinct sets.
func (uf *UnionFind) NumSets() int {
	seenRoots := NewSet()
	for i := range len(uf.parent) {
		seenRoots.Add(uf.Find(i))
	}
	return seenRoots.Len()
}
