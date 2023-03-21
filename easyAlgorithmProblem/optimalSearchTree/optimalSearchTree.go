package optimalSearchTree

import (
	"math"
	"workSpace/convenientMethod"
)

func OptimalSearchTree(names []string, probabilities []float64) (sumCount float64) {
	w := convenientMethod.Create2DSlice[float64](len(names) + 1)
	m := convenientMethod.Create2DSlice[float64](len(names) + 2)
	for i := 1; i <= len(names); i++ {
		for j := i; j <= len(names); j++ {
			w[i][j] = 0
			for k := (i - 1) * 2; k <= j*2; k++ {
				w[i][j] += probabilities[k]
			}
		}
	}
	for i := 1; i <= len(names); i++ { //对m[i][i]进行计算
		m[i][i] = probabilities[i*2-1] + probabilities[i*2-2] + probabilities[i*2]
	}
	for i := 1; i < len(names); i++ { //一共len层，除了第一层外还剩len-1层
		for j := 1; j <= len(names)-i; j++ { //计算第i层的len-i个m
			min := math.MaxFloat64
			var temp float64 = 0
			for k := j - 1; k < j+i; k++ {
				temp = m[j][k] + m[k+2][j+i] + w[j][j+i]
				if temp < min {
					min = temp
				}
			}
			m[j][j+i] = min
		}
	}

	return m[1][len(names)]
}
