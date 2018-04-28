package main

import (
	"fmt"
	"math"
)

const Epsilon float64 = 1e-12

// Radians converts degrees to radians
func Radians(degrees float64) float64 {
	return degrees / 180. * math.Pi
}

// Degrees converts radians to degrees
func Degrees(radians float64) float64 {
	return radians / math.Pi * 180.
}

// Search is binary search similar to sort.Search but operating on float64
func Search(from, to float64, f func(float64) bool) float64 {
	m := (from + to) / 2
	for math.Abs(from-to) > Epsilon {
		if f(m) {
			to = m
		} else {
			from = m
		}
		m = (from + to) / 2
	}
	return m
}

func QuadraticEquation(a, b, c float64) (float64, float64, error) {
	if a == 0 {
		return 0, 0, fmt.Errorf("Not quadratic")
	}
	D := b*b - 4*a*c
	if D < 0 {
		return 0, 0, fmt.Errorf("Discriminant less than 0")
	}

	d := math.Sqrt(D)

	return (-b + d) / (2 * a), (-b - d) / (2 * a), nil
}
