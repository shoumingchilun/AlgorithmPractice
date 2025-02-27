package minSubArrayLen

/*
https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 209. 长度最小的子数组*/

func minSubArrayLen(target int, nums []int) int {
	minLength := len(nums) + 1
	leftIndex, rightIndex := 0, 0
	tempSum := nums[0]
	for rightIndex < len(nums) {
		if tempSum < target {
			rightIndex++
			if rightIndex < len(nums) {
				tempSum += nums[rightIndex]
			}
		} else if tempSum >= target {
			if rightIndex-leftIndex+1 < minLength {
				minLength = rightIndex - leftIndex + 1
			}
			tempSum -= nums[leftIndex]
			leftIndex++
		}
	}
	if minLength == len(nums)+1 {
		return 0
	} else {
		return minLength
	}
}
