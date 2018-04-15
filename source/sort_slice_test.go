package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSortSliceInt(t *testing.T) {
	slice := []int{4, 8, 3, 5}
	sorted := []int{3, 4, 5, 8}
	sort_Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	for i := range slice {
		if slice[i] != sorted[i] {
			t.Errorf("%d: got %v want %v", i, slice[i], sorted[i])
		}
	}
}

func TestSortSliceString(t *testing.T) {
	slice := []string{"d12", "zzzzzzzzzz", "b", "efgh"}
	sorted := []string{"b", "d12", "efgh", "zzzzzzzzzz"}
	sort_Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	for i := range slice {
		if slice[i] != sorted[i] {
			t.Errorf("%d: got %v want %v", i, slice[i], sorted[i])
		}
	}
}

func TestSortSliceSliceInt(t *testing.T) {
	slice := [][]int{[]int{1, 4}, []int{1, 8}, []int{1, 3}, []int{1, 5}}
	sorted := [][]int{[]int{1, 3}, []int{1, 4}, []int{1, 5}, []int{1, 8}}
	sort_Slice(slice, func(i, j int) bool { return slice[i][1] < slice[j][1] })
	for i := range slice {
		if slice[i][1] != sorted[i][1] {
			t.Errorf("%d: got %v want %v", i, slice[i], sorted[i])
		}
	}
}

func BenchmarkSortSliceInt(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = rand.Int()
	}
	b.ResetTimer()
	sort_Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
}

func BenchmarkSortSliceInt1(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = rand.Int()
	}
	b.ResetTimer()
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
}

func BenchmarkSortInts(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = rand.Int()
	}
	b.ResetTimer()
	sort.Ints(slice)
}

func BenchmarkSortSliceSliceInt(b *testing.B) {
	slice := make([][]int, b.N)
	for i := range slice {
		slice[i] = []int{1, rand.Int()}
	}
	b.ResetTimer()
	sort_Slice(slice, func(i, j int) bool { return slice[i][1] < slice[j][1] })
}

func BenchmarkSortSliceSliceInt1(b *testing.B) {
	slice := make([][]int, b.N)
	for i := range slice {
		slice[i] = []int{1, rand.Int()}
	}
	b.ResetTimer()
	sort.Slice(slice, func(i, j int) bool { return slice[i][1] < slice[j][1] })
}

type benchSliceSliceIntAsc [][]int

func (b benchSliceSliceIntAsc) Len() int           { return len(b) }
func (b benchSliceSliceIntAsc) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b benchSliceSliceIntAsc) Less(i, j int) bool { return b[i][1] < b[j][1] }

func BenchmarkSortSort(b *testing.B) {
	slice := make([][]int, b.N)
	for i := range slice {
		slice[i] = []int{1, rand.Int()}
	}
	b.ResetTimer()
	sort.Sort(benchSliceSliceIntAsc(slice))
}
