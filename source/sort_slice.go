package main

import (
	"reflect"
	"sort"
	"unsafe"
)

//#include <string.h>
import "C"

type sort_Funcs struct {
	LenF  func() int
	LessF func(i, j int) bool
	SwapF func(i, j int)
}

func (sf sort_Funcs) Len() int {
	return sf.LenF()
}

func (sf sort_Funcs) Less(i, j int) bool {
	return sf.LessF(i, j)
}

func (sf sort_Funcs) Swap(i, j int) {
	sf.SwapF(i, j)
}

func sort_Swapper(slice interface{}) func(i, j int) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(&reflect.ValueError{Method: "Swapper", Kind: v.Kind()})
	}

	ln := v.Len()

	// Fast path for slices of size 0 and 1. Nothing to swap.
	switch ln {
	case 0:
		return func(i, j int) { panic("reflect: slice index out of range") }
	case 1:
		return func(i, j int) {
			if i != 0 || j != 0 {
				panic("reflect: slice index out of range")
			}
		}
	}

	size := v.Index(0).Type().Size()
	if size <= 8 {
		switch size {
		case 8:
			is := []int64{}
			isHeader := (*reflect.SliceHeader)(unsafe.Pointer(&is))
			isHeader.Data = v.Pointer()
			isHeader.Len = ln
			isHeader.Cap = ln
			return func(i, j int) { is[i], is[j] = is[j], is[i] }
		case 4:
			is := []int32{}
			isHeader := (*reflect.SliceHeader)(unsafe.Pointer(&is))
			isHeader.Data = v.Pointer()
			isHeader.Len = ln
			isHeader.Cap = ln
			return func(i, j int) { is[i], is[j] = is[j], is[i] }
		case 2:
			is := []int16{}
			isHeader := (*reflect.SliceHeader)(unsafe.Pointer(&is))
			isHeader.Data = v.Pointer()
			isHeader.Len = ln
			isHeader.Cap = ln
			return func(i, j int) { is[i], is[j] = is[j], is[i] }
		case 1:
			is := []int8{}
			isHeader := (*reflect.SliceHeader)(unsafe.Pointer(&is))
			isHeader.Data = v.Pointer()
			isHeader.Len = ln
			isHeader.Cap = ln
			return func(i, j int) { is[i], is[j] = is[j], is[i] }
		}
	}

	tmp := make([]byte, size)
	iSlice := []byte{}
	iHeader := (*reflect.SliceHeader)(unsafe.Pointer(&iSlice))
	iHeader.Len = int(size)
	iHeader.Cap = int(size)
	jSlice := []byte{}
	jHeader := (*reflect.SliceHeader)(unsafe.Pointer(&jSlice))
	jHeader.Len = int(size)
	jHeader.Cap = int(size)
	return func(i, j int) {
		iHeader.Data = v.Pointer() + uintptr(i)*size
		jHeader.Data = v.Pointer() + uintptr(j)*size
		copy(tmp, iSlice)
		copy(iSlice, jSlice)
		copy(jSlice, tmp)
	}
}

func sort_Slice(slice interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(slice)
	swap := sort_Swapper(slice)
	length := rv.Len()
	sort.Sort(sort_Funcs{func() int { return length }, less, swap})
}
