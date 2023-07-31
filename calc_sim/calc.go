package calc_sim

func Calc(x, y [3]float64) float64 {
	var similarity float64 = 0.0
	for i := 0; i < 3; i++ {
		similarity += x[i] * y[i]
	}
	return similarity
}