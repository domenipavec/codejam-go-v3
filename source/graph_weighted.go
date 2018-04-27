package main

import (
	"fmt"
	"math"
)

type GraphWeighted struct {
	Graph
	Weights [][]float64
}

func NewGraphWeighted(N int) GraphWeighted {
	return GraphWeighted{
		Graph:   NewGraph(N),
		Weights: make([][]float64, N),
	}
}

func (g GraphWeighted) String() string {
	output := ""
	for i := range g.Edges {
		if i != 0 {
			output += "\n"
		}
		output += fmt.Sprintf("%d:", i)
		for j, edge := range g.Edges[i] {
			output += fmt.Sprintf(" %d (%f)", edge, g.Weights[i][j])
		}
	}
	return output
}

func (g GraphWeighted) AddEdge(u, v int, w float64) {
	g.Graph.AddEdge(u, v)
	g.Weights[u] = append(g.Weights[u], w)
}

func (g GraphWeighted) AddBiEdge(u, v int, w float64) {
	g.AddEdge(u, v, w)
	g.AddEdge(v, u, w)
}

// DjikstraArray uses array for queue and is O(N^2)
func (g GraphWeighted) DjikstraArray(source int) []float64 {
	visited := make([]bool, g.N)

	distances := make([]float64, g.N)
	for v := range distances {
		if v == source {
			continue
		}
		distances[v] = math.Inf(1)
	}

	current := source
	for {
		for j, edge := range g.Edges[current] {
			if visited[edge] {
				continue
			}
			distance := distances[current] + g.Weights[current][j]
			if distances[edge] > distance {
				distances[edge] = distance
			}
		}
		visited[current] = true

		imin := -1
		for i := range distances {
			if visited[i] {
				continue
			}
			if imin == -1 || distances[i] < distances[imin] {
				imin = i
			}
		}

		if imin == -1 || distances[imin] == math.Inf(1) {
			break
		}

		current = imin
	}

	return distances
}
