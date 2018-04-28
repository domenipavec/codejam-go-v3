package main

import (
	"math"
)

const (
	int_MAX = math.MaxInt64
	int_MIN = math.MinInt64
)

func int_Min(as ...int) int {
	min := as[0]
	for _, a := range as[1:] {
		if a < min {
			min = a
		}
	}
	return min
}

func int_Max(as ...int) int {
	max := as[0]
	for _, a := range as[1:] {
		if a > max {
			max = a
		}
	}
	return max
}

func int_Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func int_Sum(vs ...int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func int_CeilDiv(a, b int) int {
	if a == 0 {
		return 0
	}
	return ((a - 1) / b) + 1
}

func int_gcd2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func int_Gcd(as ...int) int {
	gcd := as[0]
	for _, a := range as[1:] {
		gcd = int_gcd2(gcd, a)
	}
	return gcd
}

func int_lcm2(a, b int) int {
	return a * b / int_gcd2(a, b)
}

func int_Lcm(as ...int) int {
	lcm := as[0]
	for _, a := range as[1:] {
		lcm = int_lcm2(lcm, a)
	}
	return lcm
}

func int_Pow(a, b int) int {
	if b == 0 {
		return 1
	} else if b == 1 {
		return a
	} else if b%2 == 0 {
		return int_Pow(a*a, b/2)
	} else {
		return a * int_Pow(a*a, b/2)
	}
}

func int_Pow2(a int) int {
	return 1 << uint(a)
}

func int_Log10(a int) int {
	return int(math.Log10(float64(a)))
}

func int_Log2(a int) int {
	return int_Len(uint(a)) - 1
}

func int_Round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

func int_Floor(f float64) int {
	return int(math.Floor(f))
}

func int_Ceil(f float64) int {
	return int(math.Ceil(f))
}

func int_Len(x uint) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	if x >= 1<<4 {
		x >>= 4
		n += 4
	}
	if x >= 1<<2 {
		x >>= 2
		n += 2
	}
	if x >= 1<<1 {
		x >>= 1
		n += 1
	}
	if x >= 1 {
		n += 1
	}
	return n
}
