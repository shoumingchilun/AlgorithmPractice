package LCS

import "workSpace/convenientMethod"

// LCS 最长公共子序列问题，输入两个字符序列，输出两者的最长公共子序列
func LCS(x []string, y []string) (z []string) {
	Lx := len(x)
	Ly := len(y)
	m := convenientMethod.Create2DSlice2[int](Lx+1, Ly+1)
	n := convenientMethod.Create2DSlice2[[]string](Lx+1, Ly+1)
	for i := 0; i <= Lx; i++ {
		m[i][0] = 0
		n[i][0] = []string{}
	}
	for i := 0; i <= Ly; i++ {
		m[0][i] = 0
		n[0][i] = []string{}
	}
	for i := 1; i <= Lx; i++ {
		for j := 1; j <= Ly; j++ {
			if x[i-1] == y[j-1] {
				m[i][j] = m[i-1][j-1] + 1
				n[i][j] = append(n[i-1][j-1], x[i-1])
			} else if m[i-1][j] >= m[i][j-1] {
				m[i][j] = m[i-1][j]
				n[i][j] = n[i-1][j]
			} else {
				m[i][j] = m[i][j-1]
				n[i][j] = n[i][j-1]
			}
		}
	}
	z = n[Lx][Ly]
	return
}
