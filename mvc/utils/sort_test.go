package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	s := []int{9, 8, 7, 6, 5}
	Sort(s)
	assert.NotNil(t, s)
	assert.EqualValues(t, s[0], 5)
	assert.EqualValues(t, s[len(s)-1], 9)
}

func TestBubbleSortBestCase(t *testing.T) {
	s := []int{5, 6, 7, 8, 9}
	Sort(s)
	assert.NotNil(t, s)
	assert.EqualValues(t, s[0], 5)
	assert.EqualValues(t, s[len(s)-1], 9)
}

func getTestElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	actual := getTestElements(5)
	assert.NotNil(t, actual)
	assert.EqualValues(t, 5, len(actual))
	assert.EqualValues(t, 4, actual[0])
	assert.EqualValues(t, 3, actual[1])
	assert.EqualValues(t, 2, actual[2])
	assert.EqualValues(t, 1, actual[3])
	assert.EqualValues(t, 0, actual[4])
}

func BenchmarkBubbleSort(b *testing.B) {
	values := getTestElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}
	// BenchmarkBubbleSort-4   	 1997703	       585 ns/op 1000 els
	// BenchmarkBubbleSort-4   	       1	9899500285 ns/op 100000 els
}

func BenchmarkGoSort(b *testing.B) {
	values := getTestElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(values)
	}
	// BenchmarkGoSort-4   	   23439	     44490 ns/op 1000 els
	// BenchmarkGoSort-4   	     146	   7396789 ns/op 100000 els
}

func BenchmarkSort(b *testing.B) {
	values := getTestElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(values)
	}
	// BenchmarkGoSort-4   	   23439	     44490 ns/op 1000 els
	// BenchmarkGoSort-4   	     146	   7396789 ns/op 100000 els
}
