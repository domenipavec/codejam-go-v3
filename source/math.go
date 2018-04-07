package main

import "math"

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
