package maximumContiguousFieldSum

// MaximumContiguousFieldSum 最大子段和，输入一个int类型的数组，返回其中和最大的连续子段的起始位置和结束位置
//
// 使用动态规划算法
func MaximumContiguousFieldSum(nums []int) (max, begin, end int) {
	length := len(nums)
	m := make([]int, length)
	max = 0
	b := make([]int, length)
	e := -1
	if nums[0] > 0 {
		m[0] = nums[0]
		max = m[0]
		b[0] = 0
		e = 0
	} else {
		m[0] = 0
	}
	for i := 1; i < length; i++ {
		if m[i-1] > 0 {
			m[i] = m[i-1] + nums[i]
			b[i] = b[i-1]
		} else {
			m[i] = nums[i]
			b[i] = i
		}
		if max < m[i] {
			max = m[i]
			e = i
			begin = b[i]
		}
	}

	return max, begin, e
}
