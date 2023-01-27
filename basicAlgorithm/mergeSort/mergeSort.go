package mergeSort

// MergeSort 降序排列
func MergeSort(numbers []int) (orderedNumbers []int) {
	if len(numbers) <= 1 {
		return numbers
	} else {
		mid := len(numbers) / 2
		left := MergeSort(numbers[:mid])
		right := MergeSort(numbers[mid:])
		orderedNumbers = Merge(left, right)
		return
	}
}

func Merge(left []int, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for j < len(left) && k < len(right) {
		if left[j] >= right[k] {
			result[i] = left[j]
			i++
			j++
		} else {
			result[i] = right[k]
			i++
			k++
		}
	}
	for j < len(left) {
		result[i] = left[j]
		j++
		i++
	}
	for k < len(right) {
		result[i] = right[k]
		k++
		i++
	}
	return result
}
