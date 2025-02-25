package combinationSum

/*
https://leetcode.cn/problems/combination-sum/?envType=study-plan-v2&envId=top-100-liked
39. 组合总和
*/
func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	combinationSumWithUsed([]int{}, 0, candidates, target, &ret)
	return ret
}

func combinationSumWithUsed(usedNums []int, currentSum int, candidates []int, target int, ret *[][]int) {
	for i, num := range candidates {
		if currentSum+num < target {
			combinationSumWithUsed(append(usedNums, num), currentSum+num, candidates[i:], target, ret)
		} else if currentSum+num == target {
			nums := make([]int, len(usedNums)+1, len(usedNums)+1)
			copy(nums, append(usedNums, num))
			*ret = append(*ret, nums)
			//另一种写法
			//*ret = append(*ret, append([]int(nil), append(usedNums, num)...))
		}
	}
}

// 存在一个神奇的bug，append(*ret,[]int{3,3,3,3,3,3})后，下一层调用append(usedNums, num)，可能会存在usedNums地址复用导致ret中的[]int{3,3,3,3,3,3}变成[]int{3,3,3,3,3,2}
func combinationSumWithUsed2(usedNums []int, currentSum int, candidates []int, target int, ret *[][]int) {
	for i, num := range candidates {
		if currentSum+num < target {
			combinationSumWithUsed2(append(usedNums, num), currentSum+num, candidates[i:], target, ret)
		} else if currentSum+num == target {
			*ret = append(*ret, append(usedNums, num))
		}
	}
}
