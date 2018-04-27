package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Graph struct {
	N     int
	Edges [][]int
}

func NewGraph(N int) Graph {
	return Graph{
		N:     N,
		Edges: make([][]int, N),
	}
}

func (g Graph) String() string {
	output := ""
	for i := range g.Edges {
		if i != 0 {
			output += "\n"
		}
		output += fmt.Sprintf("%d:", i)
		for _, edge := range g.Edges[i] {
			output += fmt.Sprintf(" %d", edge)
		}
	}
	return output
}

func (g Graph) AddEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
}

func (g Graph) AddBiEdge(u, v int) {
	g.AddEdge(u, v)
	g.AddEdge(v, u)
}

func (g Graph) InDegrees() []int {
	inDegrees := make([]int, g.N)
	for _, edges := range g.Edges {
		for _, v := range edges {
			inDegrees[v]++
		}
	}
	return inDegrees
}

func (g Graph) TopologicalSort() ([]int, error) {
	order := make([]int, 0, g.N)
	q := Queue{}
	visited := 0

	inDegrees := g.InDegrees()
	for i := range inDegrees {
		if inDegrees[i] == 0 {
			q.Push(i)
		}
	}

	for q.Len() > 0 {
		v := q.Pop()
		visited += 1
		order = append(order, v)

		for _, edge := range g.Edges[v] {
			inDegrees[edge]--
			if inDegrees[edge] == 0 {
				q.Push(edge)
			}
		}
	}

	if visited != g.N {
		return nil, errors.New("Topological sort not possible")
	}

	return order, nil
}

func (g Graph) MinTopologicalSort() ([]int, error) {
	order := make([]int, 0, g.N)
	q := MinHeapInt{}
	visited := 0

	inDegrees := g.InDegrees()
	for i := range inDegrees {
		if inDegrees[i] == 0 {
			heap.Push(&q, i)
		}
	}

	for len(q) > 0 {
		v := heap.Pop(&q).(int)
		visited += 1
		order = append(order, v)

		for _, edge := range g.Edges[v] {
			inDegrees[edge]--
			if inDegrees[edge] == 0 {
				heap.Push(&q, edge)
			}
		}
	}

	if visited != g.N {
		return nil, errors.New("Topological sort not possible")
	}

	return order, nil
}
