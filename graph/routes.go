package graph

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/fighterlyt/permutation"
)

type graph struct {
	edges [][]float64
}

// Routable is the set of edges between the vertices of the graph.
type Routable interface {
	Paths() ([][]int, error)
	Cost(path []int) float64
}

// Paths calculate a matrix with all possibilities of routes.
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
		data := make([]float64, len(matrix[i]))

		for j, cell := range row {
			data[j], err = strconv.ParseFloat(cell, 64)
			if err != nil {
				return nil, err
			}
		}

		g.edges = append(g.edges, data)
	}

	return &g, nil
}

// Cost returns the total length between all cities in a route.
func (g *graph) Cost(path []int) float64 {
	cost := 0.0
	length := len(path)
	origin := path[0]

	distance := func(start, end int) float64 {
		return g.edges[start][end]
	}

	for i := 0; i < length; i++ {
		if (i + 1) == length {
			cost = cost + distance(path[i], origin)
		} else {
			cost = cost + distance(path[i], path[i+1])
		}
	}

	return cost
}
