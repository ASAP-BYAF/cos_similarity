package main

import(
  "fmt"

  "github.com/ASAP-BYAF/cos_similarity/calc_sim"
  "github.com/ASAP-BYAF/cos_similarity/sort"
)

func main() {
  s1 := []float64{0.3, 0.4, 0.5, 0.9, 1.0}
  s2 := []float64{0.4, 0.1, 0.6, 1.0, 1.0}
  sim := calc_sim.Calc(s1, s2)

  fmt.Printf("sim = %g\n", sim)

  sort.Sort(s1)
  fmt.Printf("s1_sorted = %v\n", s1)
}
