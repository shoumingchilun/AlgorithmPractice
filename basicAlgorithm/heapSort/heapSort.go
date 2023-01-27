package heapSort

// HeapSort 堆排序，降序排序，提供一个无序数组，返回一个有序数组
func HeapSort(numbers []int) (orderedNumbers []int) {
	length := len(numbers)
	orderedNumbers = make([]int, length)
	for i := 0; i < length; i++ {
		orderedNumbers[i] = numbers[i]
	}
	//先建堆
	heapify(orderedNumbers)
	//再进行由0节点开始的heapify2
	for i := length - 1; i >= 0; i-- {
		orderedNumbers[i], orderedNumbers[0] = orderedNumbers[0], orderedNumbers[i]
		heapify2(orderedNumbers, i, 0)
	}
	return
}

// heapify 建堆函数，除了从(length-1)/2到0节点的每个三角形要测试一下之外，还有从上到下的一次总层数的交换。
func heapify(numbers []int) {
	length := len(numbers)
	//i为开始节点
	i := (length - 1) / 2
	for ; i >= 0; i-- {
		heapify2(numbers, length, i)
	}
}

// heapify 第n个节点的排序，只进行那个三角形和三角形往下的子树的建堆
func heapify2(numbers []int, length int, n int) {
	min := n
	left := 2*n + 1
	right := 2*n + 2
	if left < length && numbers[min] > numbers[left] {
		min = left
	}
	if right < length && numbers[min] > numbers[right] {
		min = right
	}
	if min != n {
		numbers[n], numbers[min] = numbers[min], numbers[n]
		heapify2(numbers, length, min)
	}
}
