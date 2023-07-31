package calc_sim

func Calc(x, y []float64) float64 {
	var similarity float64 = 0.0
	for i := 0; i < len(x); i++ {
		similarity += x[i] * y[i]
	}
	return similarity
}