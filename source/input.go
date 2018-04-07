package main

import (
	"bufio"
	"io"
	"log"
	"math/big"
	"strconv"
)

type Input struct {
	scanner *bufio.Scanner
}

func NewInput(i io.Reader) *Input {
	scanner := bufio.NewScanner(i)
	scanner.Split(bufio.ScanWords)
	return &Input{
		scanner: scanner,
	}
}

func (i *Input) Scan() {
	if ok := i.scanner.Scan(); !ok {
		log.Fatalln("Error scanning input:", i.scanner.Err())
	}
}

func (i *Input) String() string {
	i.Scan()
	return i.scanner.Text()
}

func (i *Input) Bytes() []byte {
	i.Scan()
	data := i.scanner.Bytes()
	data_copy := make([]byte, len(data))
	copy(data_copy, data)
	return data_copy
}

func (i *Input) Int() int {
	n, err := strconv.Atoi(i.String())
	if err != nil {
		log.Fatalln("Error scanning for int:", err)
	}
	return n
}

func (i *Input) Float64() float64 {
	f, err := strconv.ParseFloat(i.String(), 64)
	if err != nil {
		log.Fatalln("Error scanning for float:", err)
	}
	return f
}

func (i *Input) BigInt() *big.Int {
	n := &big.Int{}
	str := i.String()

	n, ok := n.SetString(str, 10)
	if !ok {
		log.Fatalln("Error scanning for big int:", str)
	}
	return n
}

func (i *Input) Digits() []int {
	str := i.String()
	ints := make([]int, 0, len(str))
	for _, chr := range str {
		if chr < 48 || chr > 57 {
			log.Fatalln("String element not a digit:", chr)
		}
		ints = append(ints, int(chr-48))
	}
	return ints
}

func (i *Input) SliceInt(n int) []int {
	newSlice := make([]int, n)
	for j := 0; j < n; j++ {
		newSlice[j] = i.Int()
	}
	return newSlice
}

func (i *Input) SliceFloat64(n int) []float64 {
	newSlice := make([]float64, n)
	for j := 0; j < n; j++ {
		newSlice[j] = i.Float64()
	}
	return newSlice
}

func (i *Input) SliceString(n int) []string {
	strs := make([]string, 0, n)
	for j := 0; j < n; j++ {
		strs = append(strs, i.String())
	}
	return strs
}

func (i *Input) SliceBytes(n int) [][]byte {
	sb := make([][]byte, n)
	for j := range sb {
		sb[j] = i.Bytes()
	}
	return sb
}

func (input *Input) SliceSliceInt(n, m int) [][]int {
	newSlice := make([][]int, n)
	for i := 0; i < n; i++ {
		newSlice[i] = make([]int, m)
		for j := 0; j < m; j++ {
			newSlice[i][j] = input.Int()
		}
	}
	return newSlice
}

func (input *Input) SliceSliceFloat64(n, m int) [][]float64 {
	newSlice := make([][]float64, n)
	for i := 0; i < n; i++ {
		newSlice[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			newSlice[i][j] = input.Float64()
		}
	}
	return newSlice
}

func (input *Input) SliceSliceString(n, m int) [][]string {
	newSlice := make([][]string, n)
	for i := 0; i < n; i++ {
		newSlice[i] = make([]string, m)
		for j := 0; j < m; j++ {
			newSlice[i][j] = input.String()
		}
	}
	return newSlice
}
