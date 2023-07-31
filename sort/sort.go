package sort

import (
	"sort"
)

func RevSort(s1 []float64) []float64{

	// 降順
	sort.Sort(sort.Reverse(sort.Float64Slice(s1)))
	return s1
}