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

// RoundRobinSchedule nmd做不出来
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
		//直接计算
		m++
		roundRobinSchedule2(m, a) //获得四分之一多一点

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
	////先获得n/2是奇数还是偶数，如果是奇数，需要特别计算，如果是偶数，可降规模，然后重新计算
	//if n&1 != 1 {
	//	m := n / 2
	//	roundRobinSchedule2(m, a)
	//	for i := 0; i < m; i++ {
	//		for j := 0; j < m; j++ {
	//			a.a[m+i][m+j] = a.a[i][j]
	//			a.a[i][m+j] = a.a[i][j] + m
	//			a.a[m+i][j] = a.a[i][j] + m
	//		}
	//	}
	//} else {
	//	n++
	//	roundRobinSchedule2(n, a) //获得n+1时的图像，到时候再获得前n行n+1列的内容；现在n为4
	//	n--                       //还原n,现在n是3
	//	//for i := 0; i < n; i++ {  //先获得下方和后方1列的未处理数据
	//	//	for j := 0; j < n+1; j++ {
	//	//		a.a[n+i][j] = a.a[i][j] + n
	//	//	}
	//	//}
	//
	//}
}
