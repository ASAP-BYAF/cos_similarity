package stdize_vec

import(
	"math"
)

func Stdize(v []float64) []float64{
	var size float64 = 0.0

	// slice のサイズを計算
	for i := 0; i < len(v); i++ {
		size += v[i] * v[i]
	}
	size = math.Pow(size, 0.5)

	// slice を規格化
	for i := 0; i < len(v); i++ {
		v[i] /= size
	}

	return v
}