package fullNumberArrange

type ret struct {
	ans []int
}

// FullNumberArrange 输入数字N，将返回由1~N自然数组成的全排列数组
func FullNumberArrange(n int) []int {
	ans := make([]int, 0)
	ret := ret{ans: ans}
	for i := 1; i <= n; i++ {
		head(n, i*pow10(n-1), &ret)
	}
	return ret.ans
}
func head(n int, temp int, ret *ret) { //n指明范围为1-N，temp包含已被使用的数字
	used, usedSlice := getUsedAndUsedSlice(temp)
	if used >= n {
		ret.ans = append(ret.ans, temp)
	} else {
		unused := getUnused(n, usedSlice)
		for i := 0; i < len(unused); i++ {
			head(n, temp+unused[i]*pow10(n-used-1), ret)
		}
	}
}
func pow10(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret *= 10
	}
	return ret
}
func getUsedAndUsedSlice(temp int) (used int, s []int) {
	for temp != 0 {
		end := temp % 10
		if end != 0 {
			used++
			s = append(s, end)
		}
		temp /= 10
	}
	return
}
func isUsed(used []int, num int) bool {
	for i := 0; i < len(used); i++ {
		if used[i] == num {
			return true
		}
	}
	return false
}
func getUnused(n int, usedSlice []int) (rst []int) {
	rst = make([]int, 0)
	for i := 1; i <= n; i++ {
		if !isUsed(usedSlice, i) {
			rst = append(rst, i)
		}
	}
	return
}
