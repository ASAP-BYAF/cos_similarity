package main

import(
  "github.com/ASAP-BYAF/cos_similarity/calc"
)

func main() {
  v1 := [3]float64{0.3, 0.4, 0.5}
  v2 := [3]float64{0.4, 0.1, 0.6}
  calc(v1, v2)
}
