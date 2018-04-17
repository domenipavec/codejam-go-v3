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

func (v twod_Vector) Normalize() twod_Vector {
	return v.Div(v.Len())
}

func (v twod_Vector) Div(b float64) twod_Vector {
	return twod_Vector{
		X: v.X / b,
		Y: v.Y / b,
	}
}

func (v twod_Vector) Mul(b float64) twod_Vector {
	return twod_Vector{
		X: v.X * b,
		Y: v.Y * b,
	}
}

func (v twod_Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v twod_Vector) SameDirection(b twod_Vector) bool {
	d := v.Dot(b)
	return d*d == v.Len2()*b.Len2() && d > 0
}

func (v twod_Vector) Parallel(b twod_Vector) bool {
	d := v.Dot(b)
	return d*d == v.Len2()*b.Len2()
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
	Data []twod_Vector
}

func (g *twod_Graph) AddXY(X, Y float64) {
	g.Add(twod_Vector{X: X, Y: Y})
}

func (g *twod_Graph) Add(v twod_Vector) {
	g.Data = append(g.Data, v)
}

func (g twod_Graph) Len() int {
	return len(g.Data)
}

func (g twod_Graph) Area(indices ...int) float64 {
	area := 0.
	if len(indices) > 0 {
		for idx, i := range indices {
			if idx < len(indices)-1 {
				area += g.Data[i].Cross(g.Data[indices[idx+1]])
			} else {
				area += g.Data[i].Cross(g.Data[indices[0]])
			}
		}
	} else {
		for i := range g.Data {
			if i < len(g.Data)-1 {
				area += g.Data[i].Cross(g.Data[i+1])
			} else {
				area += g.Data[i].Cross(g.Data[0])
			}
		}
	}

	return math.Abs(area) / 2
}

func (g twod_Graph) ConvexHull(includeCollinears ...bool) []int {
	if g.Len() < 3 {
		return nil
	}
	collinears := len(includeCollinears) > 0 && includeCollinears[0]
	hull := make([]int, 0, 10)

	l := 0
	for i := 1; i < g.Len(); i++ {
		if g.Data[i].X == g.Data[l].X {
			if g.Data[i].Y < g.Data[l].Y {
				l = i
			}
		} else if g.Data[i].X < g.Data[l].X {
			l = i
		}
	}

	p := l
	prevToP := twod_Vector{}
	for {
		hull = append(hull, p)

		q := (p + 1) % g.Len()
		pToq := g.Data[q].Sub(g.Data[p])
		for !prevToP.IsZero() && prevToP.Parallel(pToq) && !prevToP.SameDirection(pToq) {
			q = (q + 1) % g.Len()
			pToq = g.Data[q].Sub(g.Data[p])
		}
		for i := 0; i < g.Len(); i++ {
			if i == p {
				continue
			}
			pToi := g.Data[i].Sub(g.Data[p])
			qToi := g.Data[q].Sub(g.Data[i])
			cross := pToi.Cross(qToi)
			if cross == 0 {
				if (!collinears && pToi.Len2() > pToq.Len2()) ||
					(collinears && pToi.Len2() < pToq.Len2() && (prevToP.IsZero() || !prevToP.Parallel(pToi) || prevToP.SameDirection(pToi))) {
					q = i
					pToq = pToi
				}
			} else if cross > 0 {
				q = i
				pToq = pToi
			}
		}

		p = q
		prevToP = pToq

		if p == l {
			break
		}
	}

	return hull
}

func (g twod_Graph) ValidPolygon(indices ...int) bool {
	if len(indices) == 0 {
		indices = Range(g.Len())
	}

	for idx := range indices {
		p := g.Data[indices[idx]]
		var pr, nxt twod_Vector
		if idx < len(indices)-1 {
			pr = g.Data[indices[idx+1]]
			if idx < len(indices)-2 {
				nxt = g.Data[indices[idx+2]]
			} else {
				nxt = g.Data[indices[0]]
			}
		} else {
			pr = g.Data[indices[0]]
			nxt = g.Data[indices[1]]
		}

		r := pr.Sub(p)
		rn := nxt.Sub(pr)

		if r.Parallel(rn) && !r.SameDirection(rn) {
			return false
		}

		for idx1 := idx + 2; idx1 < len(indices); idx1++ {
			q := g.Data[indices[idx1]]
			var qs twod_Vector
			if idx1 < len(indices)-1 {
				qs = g.Data[indices[idx1+1]]
			} else {
				if idx == 0 {
					continue
				}
				qs = g.Data[indices[0]]
			}

			if !LineIntersection(p, pr, q, qs).IsZero() {
				return false
			}
		}
	}

	return true
}

// LineIntersection lines from p to pr and from q to qs
func LineIntersection(p, pr, q, qs twod_Vector) twod_Vector {
	r := pr.Sub(p)
	s := qs.Sub(q)

	rs := r.Cross(s)
	qmp := q.Sub(p)
	qpr := qmp.Cross(r)

	if rs == 0 && qpr == 0 {
		r2 := r.Len2()
		t0 := qmp.Dot(r) / r2
		t1 := qmp.Add(s).Dot(r) / r2
		if t0 <= 1 && t0 >= 0 {
			return p.Add(r.Mul(t0))
		}
		if t1 <= 1 && t1 >= 0 {
			return p.Add(r.Mul(t1))
		}
		if (t1 >= 1 && t0 <= 0) || (t1 <= 0 && t0 >= 1) {
			return p
		}
	} else if rs != 0 {
		u := qpr / rs
		t := qmp.Cross(s) / rs
		if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
			return p.Add(r.Mul(t))
		}
	}
	return twod_Vector{}
}
