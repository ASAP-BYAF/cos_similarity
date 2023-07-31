package calc_sim

import(
	"github.com/ASAP-BYAF/cos_similarity/calc_sim/stdize_vec"
)
func Calc(x, y []float64) float64 {

	// 規格化
	stdize_vec.Stdize(x)
	stdize_vec.Stdize(y)
  
	var similarity float64 = 0.0
	for i := 0; i < len(x); i++ {
		similarity += x[i] * y[i]
	}
	return similarity
}