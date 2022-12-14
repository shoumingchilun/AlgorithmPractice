package convenientMethod

// Create2DSlice 生产一个类型为T的n*n的2维空切片
func Create2DSlice[T any](n int) [][]T {
	a := make([][]T, n)
	for i := 0; i < n; i++ {
		a[i] = make([]T, n)
	}
	return a
}

// Create2DSlice2 创建一个大小为n*m的二维切片
func Create2DSlice2[T any](n int, m int) [][]T {
	a := make([][]T, n)
	for i := 0; i < n; i++ {
		a[i] = make([]T, m)
	}
	return a
}
