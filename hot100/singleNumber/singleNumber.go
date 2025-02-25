package singleNumber

/*
https://leetcode.cn/problems/single-number/submissions/603110112/?envType=study-plan-v2&envId=top-100-liked
136. 只出现一次的数字
*/

func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret
}
