package multipleMachineScheduling

import (
	"fmt"
)

type task struct {
	t     int //用来记录任务时长
	index int //用来记录任务号
}

// MultipleMachineScheduling 输入ti为各个任务需要时间，m为机器数；输出为需要总时间；会将调度方案打印
func MultipleMachineScheduling(ti []int, m int) (all int) {
	if m == 1 {
		fmt.Print("机器1：")
		for i := 0; i < len(ti); i++ {
			fmt.Print("任务", i, ",")
			all += ti[i]
		}
		fmt.Println("耗时：", all)
		return all
	}
	if m >= len(ti) {
		for i := 0; i < len(ti); i++ {
			if ti[i] > all {
				all = ti[i]
			}
			fmt.Println("机器", i, ":任务", i)
		}
		fmt.Println("耗时：", all)
		return all
	}
	//下面才是正常的逻辑
	//1.将[]int进行转换成[]task
	tasks := make([]task, len(ti))
	for i := 0; i < len(ti); i++ {
		tasks[i] = task{ti[i], i}
	}
	//2.对task进行排序
	tasks = QuickSort(tasks)
	//3.开始调度逻辑
	ts := make([]int, m)   //时间戳，用来记录各机器下次空闲时间的时间戳
	s := make([]string, m) //记录任务顺序
	for i := 0; i < m; i++ {
		s[i] = fmt.Sprint("机器", i, "：")
	}
	//4.分配任务
	for i := 0; i < len(tasks); i++ {
		index := 0
		minTs := ts[0]
		for j := 1; j < m; j++ {
			if ts[j] < minTs {
				index = j
				minTs = ts[j]
			}
		}
		ts[index] += tasks[i].t
		s[index] = fmt.Sprint(s[index], "任务", i, ",")
	}
	for i := 0; i < m; i++ {
		s[i] = fmt.Sprint(s[i], "耗时：", ts[i])
		fmt.Println(s[i])
	}
	all = ts[0]
	for i := 0; i < m; i++ {
		if ts[i] > all {
			all = ts[i]
		}
	}
	fmt.Println("总耗时：", all)
	return all
}

// QuickSort 降序排列
func QuickSort(numbers []task) (orderedNumbers []task) {
	if len(numbers) <= 1 {
		return numbers[:]
	}
	flag := true
	i, j, k := 0, 0, len(numbers)-1
	for j != k {
		if numbers[k].t > numbers[i].t && flag {
			numbers[k], numbers[i] = numbers[i], numbers[k]
			i = k
			j++
			flag = false
			continue
		} else if numbers[j].t < numbers[i].t && !flag {
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
