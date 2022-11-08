package quickSort

// QuickSort 降序排列
func QuickSort(numbers []int) (orderedNumbers []int) {
	if len(numbers) <= 1 {
		return numbers[:]
	}
	flag := true
	i, j, k := 0, 0, len(numbers)-1
	for j != k {
		if numbers[k] > numbers[i] && flag {
			numbers[k], numbers[i] = numbers[i], numbers[k]
			i = k
			j++
			flag = false
			continue
		} else if numbers[j] < numbers[i] && !flag {
			numbers[j], numbers[i] = numbers[i], numbers[j]
			i = j
			k--
			flag = true
			continue
		}
		if flag {
			k--
		} else {
			j++
		}
	}
	left := QuickSort(numbers[:j])
	right := QuickSort(numbers[j+1:])
	left = append(left, numbers[j])
	return append(left, right...)
}
