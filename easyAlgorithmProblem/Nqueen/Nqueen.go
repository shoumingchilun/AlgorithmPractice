package Nqueen

import "fmt"

// Nqueen 输入N皇后问题的n，将返回解的数量并打印全部的解
func Nqueen(n int) int {
	cols := make([]int, n) //用来暂时记录已经放了皇后的位置
	total := queen(n, 0, cols, 0)
	return total
}

func queen(n int, row int, cols []int, total int) int {
	if row == n {
		total++
		fmt.Println(cols)
	} else {
		for i := 0; i < n; i++ {
			cols[row] = i
			if judge(row, cols) {
				total = queen(n, row+1, cols, total)
			}
		}
	}
	return total
}

func judge(row int, cols []int) bool {
	for j := 0; j != row; j++ {
		if cols[row] == cols[j] || row-cols[row] == j-cols[j] || row+cols[row] == j+cols[j] {
			return false
		}
	}
	return true
}
