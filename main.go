package main

import (
	"fmt"
	"workSpace/easyAlgorithmProblem/maximumContiguousFieldSum"
)

func main() {
	max, b, e := maximumContiguousFieldSum.MaximumContiguousFieldSum([]int{-2, 11, -4, 13, -5, -2})
	fmt.Println("最长字段和：", max, "，起始位置：", b, "，结束位置：", e, "。")
}
