package roundRobinSchedule

import (
	"fmt"
	"workSpace/convenientMethod"
)

type Schedule struct {
	a [][]int
}

// SimpleRoundRobinSchedule 队伍数量为2的n次方时的日程安排，第一列为队伍序号，从第二列开始依次是此队伍第1,2,...天的对战方序号
func SimpleRoundRobinSchedule(n int) [][]int {
	numbers := 1 << n
	var a Schedule
	a.a = convenientMethod.Create2DSlice[int](numbers)
	simpleCopy(n, &a)
	return a.a
}

func simpleCopy(n int, a *Schedule) {
	if n == 0 {
		a.a[0][0] = 1
	} else {
		simpleCopy(n-1, a)
		m := 1 << (n - 1)
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				a.a[m+i][m+j] = a.a[i][j]
				a.a[i][m+j] = a.a[i][j] + m
				a.a[m+i][j] = a.a[i][j] + m
			}
		}
	}
}

// RoundRobinSchedule nmd终于做出来了
// RoundRobinSchedule 传入队伍数量n，n为任意正偶数，返回日程安排数组int[][]类似Simple
func RoundRobinSchedule(n int) [][]int {
	if n&1 == 1 || n <= 0 {
		fmt.Println("输入错误")
	}
	var a Schedule
	a.a = convenientMethod.Create2DSlice[int](n)
	roundRobinSchedule2(n, &a)
	return a.a
}

func roundRobinSchedule2(n int, a *Schedule) {
	if n == 1 {
		a.a[0][0] = 1
		return
	}
	m := n / 2
	if m&1 == 1 && m != 1 { //说明n不可降规模，需要直接计算，m集合为3,5,7,...
		//直接计算前m+1列
		roundRobinSchedule2(m+1, a) //获得四分之一多一点
		indexes := make([]int, m)   //用于记录超出n的数的位置
		for i := 0; i < m; i++ {    //向下重复并加m-1，同时记录超出n的数的位置
			for j := 0; j < m+1; j++ {
				a.a[m+i][j] = a.a[i][j] + m
				if a.a[m+i][j] > n {
					indexes[i] = j
				}
			}
		}
		for i := 0; i < m; i++ { //更改越界的值，同时向上映射
			a.a[m+i][indexes[i]] = i + 1
			a.a[i][indexes[i]] = i + 1 + m
		}
		//计算后m-1列
		numbers := make([]int, 2*m) //numbers用于存储两遍m+1~n
		for i := 0; i < m; i++ {
			numbers[i] = m + 1 + i
			numbers[m+i] = m + 1 + i
		}
		for i := 0; i < m; i++ { //向右赋值
			for j := m + 1; j < n; j++ {
				a.a[i][j] = numbers[j-m+i]
				a.a[numbers[j-m+i]-1][j] = i + 1
			}
		}
	} else { //说明n可以继续降规模,m集合为1,2,4,6,...
		//计算四分之一后扩大到整个n
		roundRobinSchedule2(m, a)
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				a.a[m+i][m+j] = a.a[i][j]
				a.a[i][m+j] = a.a[i][j] + m
				a.a[m+i][j] = a.a[i][j] + m
			}
		}
	}
}
