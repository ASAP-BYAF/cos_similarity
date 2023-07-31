package main

import(
  "fmt"

  "github.com/ASAP-BYAF/cos_similarity/calc_sim"
  "github.com/ASAP-BYAF/cos_similarity/sort"
)

func main() {
  // s1 を質問分をベクトルしたものとする。
  s1 := []float64{0.3, 0.4, 0.5}
  s2 := []float64{0.4, 0.1, 0.6}
  s3 := []float64{0.2, 0.1, 0.5}
  s4 := []float64{1.3, 3.5, 2.5}
  s_list := [][]float64{s2, s3, s4}

  // ファイルのベクトルのそれぞれに対して質問文とのコサイン類似度を計算
  // sim_list に格納。
  var sim float64 = 0.0
  sim_list := []float64{}
  for i := 0; i < len(s_list); i++ {
    sim = calc_sim.Calc(s1, s_list[i])
    sim_list = append(sim_list, sim)
	}

  // 優先度の高い順にソート
  sort.RevSort(sim_list)

  // 確認
  fmt.Printf("sim_list_sorted = %v\n", sim_list)
}
