package roundRobinSchedule

type Schedule struct {
	a [][]int
}

// SimpleRoundRobinSchedule 队伍数量为2的n次方时的日程安排
func SimpleRoundRobinSchedule(n int) [][]int {
	numbers := 1 << n
	var a Schedule
	a.a = makeSlice(numbers)
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

// RoundRobinSchedule 传入队伍数量n，返回日程安排数组int[][]
func RoundRobinSchedule(n int) [][]int {
	var a Schedule
	a.a = makeSlice(n)
	robinSchedule2(n, &a)
	return a.a
}

func robinSchedule2(n int, a *Schedule) {
	if n == 1 {
		a.a[0][0] = 1
		return
	}
	//判断n为偶数还是奇数
	if n&1 == 1 { //奇数
		robinSchedule2(n+1, a)
	} else { //偶数
		robinSchedule2(n/2, a)
	}
	makeCopy(n, a)
}

// makeSlice 用来生成n*n的二维空切片
func makeSlice(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	return a
}

func makeCopy(n int, a *Schedule) {
	m := n / 2
	if m&1 == 1 {
		copyOdd(n, a)
	} else {
		copyN(n, a)
	}
}
func copyN(n int, b *Schedule) {
	m := n / 2
	for i := 0; i <= m-1; i++ {
		for j := 0; j <= m-1; j++ {
			b.a[i][j+m] = b.a[i][j] + m
			b.a[i+m][j] = b.a[i][j+m]
			b.a[i+m][j+m] = b.a[i][j]
		}
	}
}
func copyOdd(n int, b *Schedule) {
	m := n / 2
	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = m + i + 1
	}
	for i := 0; i <= m-1; i++ {
		for j := 0; j <= m-1; j++ {
			if b.a[i][j] > m {
				b.a[i][j] = a[i]
				b.a[i+m][j] = i + 1
			} else {
				b.a[i+m][j] = b.a[i][j] + m
			}
		}
		for j := 1; j <= m-1; j++ {
			b.a[i][m+j] = a[i+j]
			b.a[a[i+j]][m+j] = i + 1
		}
	}
}
