package combinationSum3

/*
https://leetcode.cn/problems/combination-sum-iii/description/
216. 组合总和 III
*/

func combinationSum3(k int, n int) [][]int {
	ret := [][]int{}
	combinationSum3WithCurrent(k, n, 0, make([]bool, 9), &ret, 1)
	return ret
}
func combinationSum3WithCurrent(k int, n int, sum int, used []bool, ret *[][]int, currentIndex int) {
	if sum == n && k == 0 {
		retNums := []int{}
		for i := 1; i <= 9; i++ {
			if used[i-1] {
				retNums = append(retNums, i)
			}
		}
		*ret = append(*ret, retNums)
		return
	} else if k == 0 && sum != n {
		return
	}
	if k < 0 || sum > n || currentIndex > 9 {
		return
	}
	used[currentIndex-1] = true
	combinationSum3WithCurrent(k-1, n, sum+currentIndex, used, ret, currentIndex+1)
	used[currentIndex-1] = false
	combinationSum3WithCurrent(k, n, sum, used, ret, currentIndex+1)
}
