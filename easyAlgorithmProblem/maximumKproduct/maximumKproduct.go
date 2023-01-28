package maximumKproduct

import (
	"fmt"
	"math"
	"workSpace/convenientMethod"
)

// MaximumKProduct 向一个数字串中插入k-1个乘号，将这个数分成k部分，找出一种分法，使得这k个部分的乘积最大。
// MaximumKProduct 输入一个int和k，输出最大乘积(懒得记录是那些数相乘了...想要实现的话可以自己加一个数据结构来跟踪每次m计算时的数据)
func MaximumKProduct(number int, k int) (max int) {
	length := len(fmt.Sprint(number))
	//w数组用w函数代替，直接上边界条件
	m := convenientMethod.Create2DSlice2[int](length, k)
	for i := 0; i < length; i++ {
		m[i][0] = w(number, 0, i)
	}
	//竖着一列列地构建
	for j := 1; j < k; j++ {
		m[j][j] = m[j-1][j-1] * w(number, j, j)
		for i := j + 1; i < length; i++ {
			//比较计算获得最大值
			max2 := math.MinInt
			for h := j - 1; h < i; h++ {
				if max2 < m[h][j-1]*w(number, h+1, i) {
					max2 = m[h][j-1] * w(number, h+1, i)
				}
			}
			m[i][j] = max2
		}
	}
	max = m[length-1][k-1]
	return
}

func w(number int, begin int, end int) int {
	length := len(fmt.Sprint(number))
	number = number % (int(math.Pow10(length - begin)))
	number = number / (int(math.Pow10(length - end - 1)))
	return number
}
