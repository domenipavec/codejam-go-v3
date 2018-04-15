package main

import (
	"math"
)

type twod_Vector struct {
	X float64
	Y float64
}

func (v twod_Vector) Add(b twod_Vector) twod_Vector {
	return twod_Vector{
		X: v.X + b.X,
		Y: v.Y + b.Y,
	}
}

func (v twod_Vector) Sub(b twod_Vector) twod_Vector {
	return twod_Vector{
		X: v.X - b.X,
		Y: v.Y - b.Y,
	}
}

func (v twod_Vector) Dot(b twod_Vector) float64 {
	return v.X*b.X + v.Y*b.Y
}

func (v twod_Vector) Cross(b twod_Vector) float64 {
	return v.X*b.Y - b.X*v.Y
}

func (v twod_Vector) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v twod_Vector) Len() float64 {
	return math.Sqrt(v.Len2())
}

type twod_Graph struct {
	data []twod_Vector
}

func (g *twod_Graph) AddXY(X, Y float64) {
	g.Add(twod_Vector{X: X, Y: Y})
}

func (g *twod_Graph) Add(v twod_Vector) {
	g.data = append(g.data, v)
}

func (g twod_Graph) Get(i int) twod_Vector {
	return g.data[i]
}

func (g twod_Graph) Len() int {
	return len(g.data)
}

func (g twod_Graph) ConvexHull() []int {
	if g.Len() < 3 {
		return nil
	}
	hull := make([]int, 0, 10)

	l := 0
	for i := 1; i < g.Len(); i++ {
		if g.data[i].X == g.data[l].X {
			if g.data[i].Y < g.data[l].Y {
				l = i
			}
		} else if g.data[i].X < g.data[l].X {
			l = i
		}
	}

	p := l
	for {
		hull = append(hull, p)

		q := (p + 1) % g.Len()
		for i := 0; i < g.Len(); i++ {
			cross := g.data[i].Sub(g.data[p]).Cross(g.data[q].Sub(g.data[i]))
			if cross == 0 {
				if g.data[i].Sub(g.data[p]).Len2() > g.data[q].Sub(g.data[p]).Len2() {
					q = i
				}
			} else if cross > 0 {
				q = i
			}
		}

		p = q

		if p == l {
			break
		}
	}

	return hull
}
