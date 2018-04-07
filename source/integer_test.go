package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, int_Min(1, 3))
	assert.Equal(t, 2, int_Min(15, 2))
	assert.Equal(t, 0, int_Min(0, 1))
	assert.Equal(t, -1, int_Min(1, -1))
	assert.Equal(t, 0, int_Min(5, 4, 3, 2, 1, 0))
	assert.Equal(t, -1, int_Min(5, 4, 3, -1, 1, 0))
	assert.Equal(t, 1, int_Min(1))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, int_Max(1, 3))
	assert.Equal(t, 15, int_Max(15, 2))
	assert.Equal(t, 1, int_Max(0, 1))
	assert.Equal(t, -1, int_Max(-20, -1))
	assert.Equal(t, 5, int_Max(5, 4, 3, 2, 1, 0))
	assert.Equal(t, 10, int_Max(5, 4, 3, 10, 1, 0))
	assert.Equal(t, 1, int_Max(1))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, int_Abs(1))
	assert.Equal(t, 1, int_Abs(-1))
	assert.Equal(t, 1234, int_Abs(1234))
	assert.Equal(t, 1234, int_Abs(-1234))
	assert.Equal(t, 0, int_Abs(0))
}

func TestCeilDiv(t *testing.T) {
	assert.Equal(t, 1, int_CeilDiv(5, 5))
	assert.Equal(t, 2, int_CeilDiv(6, 5))
	assert.Equal(t, 2, int_CeilDiv(8, 5))
	assert.Equal(t, 2, int_CeilDiv(10, 5))
	assert.Equal(t, 3, int_CeilDiv(11, 5))
	assert.Equal(t, 0, int_CeilDiv(0, 2))
	assert.Equal(t, 1, int_CeilDiv(1, 2))
}

func TestGcd(t *testing.T) {
	assert.Equal(t, 2, int_Gcd(40, 6, 1000))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 30, int_Lcm(2, 5, 10, 3))
}

func TestLen(t *testing.T) {
	assert.Equal(t, int_Len(1), 1)
	assert.Equal(t, int_Len(2), 2)
	assert.Equal(t, int_Len(3), 2)
	assert.Equal(t, int_Len(127), 7)
	assert.Equal(t, int_Len(65536), 17)
	assert.Equal(t, int_Len(9999999999999999999), 64)
}

func BenchmarkGcd(b *testing.B) {
	as := make([]int, b.N)
	bs := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		as[i] = rand.Int()
		bs[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		int_gcd2(as[i], bs[i])
	}
}
