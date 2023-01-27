package LIS

// LIS 给出一个序列a1,a2,a3,a4,a5,a6,a7....an，求它的一个子序列（设为s1,s2,...sn），使得这个子序列满足这样的性质，s1<s2<s3<...<sn并且这个子序列的长度最长。
// LIS 输入：int切片类型的a数组
// LIS 输出：int切片类型的s子序列和int类型的长度
func LIS(a []int) (minList []int, minLength int) {

	minLength = 1
	index := 0
	length := len(a)
	pre := make([]int, length)
	b := make([]int, length)
	for i := 0; i < length; i++ {
		b[i], pre[i] = 1, -1
	}
	for i := 1; i < length; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if a[j] < a[i] && b[j] > max {
				max = b[j]
				pre[i] = j
			}
		}
		b[i] = max + 1
		if b[i] > minLength {
			minLength = b[i]
			index = i
		}
	}

	for index != -1 {
		minList = append([]int{index}, minList...)
		index = pre[index]
	}
	return
}
