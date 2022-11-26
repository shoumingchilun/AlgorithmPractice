package chorusFormation

// chorusFormation N位同学站成一排，音乐老师要请其中的(N-K)位同学出列，使得剩下的K位同学排成合唱队形。合唱队形是指这样的一种队形：设K位同学从左到右依次编号为1, 2, …, K，他们的身高分别为T1, T2, …, TK，则他们的身高满足T1 < T2 < … < Ti , Ti > Ti+1 > … > TK (1 <= i <= K)。
// chorusFormation 已知所有N位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。（注：身高不能相等）
// chorusFormation 输入：代表身高的int切片；输出：代表保留人员序号的int切片和最小出列人数
func chorusFormation(a []int) (index []int, minOut int) {
	//思路：计算每个index左侧最多人数和右侧最多人数,然后相加
	length := len(a)

	minOut = length
	middle := 0
	b := make([]int, length)
	c := make([]int, length)
	preB := make([]int, length)
	preC := make([]int, length)

	for i := 0; i < length; i++ {
		b[i], c[i] = 1, 1
		preB[i], preC[i] = -1, length
	}
	for i := 0; i < length; i++ {
		max1 := 0
		max2 := 0
		for j := 0; j < i; j++ {
			if a[j] < a[i] && b[j] > max1 {
				max1 = b[j]
				preB[i] = j
			}
		}
		for k := length - 1; k > length-i-1; k-- {
			if a[length-i-1] > a[k] && c[k] > max2 {
				max2 = c[k]
				preC[length-i-1] = k
			}
		}
		b[i] = max1 + 1
		c[length-i-1] = max2 + 1
	}
	for i := 0; i < length; i++ {
		out := length - (b[i] + c[i] - 1)
		if out < minOut {
			minOut = out
			middle = i
		}
	}
	before, after := middle, middle
	for before != -1 {
		index = append([]int{before}, index...)
		before = preB[before]
	}
	after = preC[after]
	for after != length {
		index = append(index, after)
		after = preC[after]
	}
	return
}
