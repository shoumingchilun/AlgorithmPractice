package knapsackProblem

import (
	"workSpace/convenientMethod"
)

// knapsack01 给定n种物品和一背包。物品i的重量是wi，其价值为vi，背包的容量为C。问应如何选择装入背包的物品，使得装入背包中物品的总价值最大?
// knapsack01 输入：代表重量的int切片，代表价值的int切片，背包容量
// knapsack01 输出：代表被选择货物的int切片，最高价值
func knapsack01(weight []int, value []int, c int) (selected []int, all int) {
	number := len(weight)
	m := convenientMethod.Create2DSlice2[int](number, c+1)       //就是那张推演的表
	index := convenientMethod.Create2DSlice2[[]int](number, c+1) //记录被选中的物品
	//设置第一行和第一列的数据，作为后面递推会用到的边界条件
	for i := 0; i < number; i++ {
		m[i][0] = 0
		index[i][0] = make([]int, 1)
		index[i][0][0] = -1
	}
	for i := 0; i <= c; i++ {
		if i >= weight[0] {
			m[0][i] = value[0]
			index[0][i] = make([]int, 1)
			index[0][i][0] = 0
		} else {
			m[0][i] = 0
			index[0][i] = make([]int, 1)
			index[0][i][0] = -1
		}
	}

	for i := 1; i < number; i++ {
		for j := 1; j <= c; j++ {
			if j > weight[i] {
				var which bool
				m[i][j], which = max(m[i-1][j], m[i-1][j-weight[i]]+value[i])
				if !which {
					index[i][j] = index[i-1][j]
				} else {
					index[i][j] = append(index[i-1][j-weight[i]], i)
				}
			} else {
				m[i][j] = m[i-1][j]
				index[i][j] = index[i-1][j]
			}
		}
	}
	all = m[number-1][c]
	selected = index[number-1][c]
	return
}
func max(a1 int, a2 int) (max int, whichOne bool) {
	if a1 >= a2 {
		return a1, false
	} else {
		return a2, true
	}
}
