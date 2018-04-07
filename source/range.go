package main

import "log"

func Range(args ...int) []int {
	start := 0
	stop := 0
	step := 1
	if len(args) == 1 {
		stop = args[0]
	} else if len(args) == 2 {
		start = args[0]
		stop = args[1]
	} else if len(args) == 3 {
		start = args[0]
		stop = args[1]
		step = args[2]
	} else {
		log.Fatal("Invalid number of args for Range")
	}
	data := make([]int, 0, (stop-start+step)/step)
	for i := start; i < stop; i += step {
		data = append(data, i)
	}
	return data
}
