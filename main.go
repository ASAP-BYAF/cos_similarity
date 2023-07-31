package main

import(
  "fmt"

  "github.com/ASAP-BYAF/cos_similarity/calc_sim"
)

func main() {
  v1 := [3]float64{0.3, 0.4, 0.5}
  v2 := [3]float64{0.4, 0.1, 0.6}
  sim := calc_sim.Calc(v1, v2)

  fmt.Printf("sim = %g", sim)
}
