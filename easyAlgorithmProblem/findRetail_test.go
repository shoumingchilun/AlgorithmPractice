package easyAlgorithmProblem

import (
	"fmt"
	"testing"
)

func TestFindRetail(t *testing.T) {
	fmt.Println("输入金额：", 123, "分")
	retail := FindRetail(123)
	fmt.Println("两角五：", retail[0], "\n一角：", retail[1], "\n五分：", retail[2], "\n一分：", retail[3])
}
