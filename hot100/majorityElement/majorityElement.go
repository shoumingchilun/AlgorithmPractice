package majorityElement

/*
https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-100-liked
169. 多数元素
*/
func majorityElement(nums []int) int {
	hashMap := make(map[int]int)
	maxNum := -1
	maxCount := 0
	for _, num := range nums {
		hashMap[num]++
		if hashMap[num] > maxCount {
			maxCount = hashMap[num]
			maxNum = num
		}
	}
	return maxNum
}
func majorityElement2(nums []int) int {
	currentNum := nums[0]
	count := 1
	for _, num := range nums[1:] {
		if currentNum == num {
			count++
		}
		if currentNum != num {
			count--
			if count == 0 {
				currentNum = num
				count = 1
			}
		}
	}
	return currentNum
}
