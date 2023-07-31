package sort

import (
	"fmt"
	"sort"
)

func Sort(s1 []float64) []float64{

	// 昇順
	sort.Sort(sort.Float64Slice(s1))
	fmt.Printf("len:%d cap:%d v:%v\n", len(s1), cap(s1), s1)

	// 降順
	sort.Sort(sort.Reverse(sort.Float64Slice(s1)))
	fmt.Printf("len:%d cap:%d v:%v\n", len(s1), cap(s1), s1)

	return s1
}