package graph

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"fmt"

	"github.com/fighterlyt/permutation"
)

type graph struct {
	edges [][]int
}

// Routable is the set of edges between the vertices of the graph.
type Routable interface {
	Paths() ([][]int, error)
	Cost(path []int) int
}

func (g *graph) Paths() ([][]int, error) {
	lenght := len(g.edges[0])
	fatorial := func(v int) int {
		r := 1
		for i := r; i <= v; i++ {
			r = r * i
		}
		return r
	}

	vertices := make([]int, lenght)
	for index := range vertices {
		vertices[index] = index
	}

	perm, err := permutation.NewPerm(vertices, nil)
	if err != nil {
		return nil, err
	}

	data := make([][]int, fatorial(lenght))
	i := 0

	for j, err := perm.Next(); err == nil; j, err = perm.Next() {
		vector := j.([]int)
		data[i] = append(data[i], vector...)
		i++
	}

	return data, nil
}

// New returns an instance of Routable given the contents of a route file.
func New(content []byte) (Routable, error) {
	result := csv.NewReader(bytes.NewReader(content))

	matrix, err := result.ReadAll()
	if err != nil {
		return nil, err
	}

	g := graph{}

	for i, row := range matrix {
		data := make([]int, len(matrix[i]))

		for j, cell := range row {
			data[j], err = strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
		}

		g.edges = append(g.edges, data)
	}

	return &g, nil
}

func (g *graph) Cost(path []int) int {
	cost := 0
	length := len(path)

	fmt.Println(path)

	distance := func(start, end int) int {
		if length == end {
			return g.edges[path[start]][path[0]]
		}
		return g.edges[path[start]][path[end]]
	}

	for i := 0; i < length; i++ {
		cost = cost + distance(i, i+1)
	}

	return cost
}
