package sum

import "testing"

//Int32 testing
var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{-5, -65, -70},
	{8, 2, 10},
	{2147483647, -1, 2147483646},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) return %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

///////////////////////////////////////////
//Uint32 testing
var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{1, 1, 2}, //1+1=2
	{2147483647, 2147483648, 4294967295},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) return %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

///////////////////////////////////////////
//Uint32 testing
var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{1, 1, 2}, //1+1=2
	{2147483647, 2147483648, 4294967295},
	{1, 2, 3},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) return %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

///////////////////////////////////////////
//Uint32 testing
var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{1, 1, 2}, //1+1=2
	{2147483647, 2147483648, 4294967295},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) return %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
