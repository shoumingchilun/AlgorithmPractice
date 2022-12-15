package batchJobScheduling

import (
	"workSpace/convenientMethod"
)

// BatchJobScheduling 输入代表两个机器耗时的两个数组，数组长度极为任务个数；输出最佳完成时间的int值和执行顺序order数组
func BatchJobScheduling(mach1, mach2 []int) (b int, order []int) {
	length := len(mach1)
	datas := data{
		n:     length,
		M:     convenientMethod.Create2DSlice2[int](length+1, 2),
		x:     make([]int, length+1),
		bestX: make([]int, length+1),
		f2:    make([]int, length+1),
		bestF: 10000,
	}
	for i := 1; i <= datas.n; i++ {
		datas.x[i] = i
		datas.M[i][0] = mach1[i-1]
		datas.M[i][1] = mach2[i-1]
	}
	backTrack(1, &datas)
	return datas.bestF, datas.bestX[1:]

}
func backTrack(i int, datas *data) {
	if i > datas.n {
		if datas.f < datas.bestF {
			for j := 1; j <= datas.n; j++ {
				datas.bestX[j] = datas.x[j]
			}
			datas.bestF = datas.f
		}
	} else {
		for j := i; j <= datas.n; j++ {
			datas.f1 += datas.M[datas.x[j]][0]
			if datas.f2[i-1] > datas.f1 {
				datas.f2[i] = datas.f2[i-1] + datas.M[datas.x[j]][1]
			} else {
				datas.f2[i] = datas.f1 + datas.M[datas.x[j]][1]
			}
			datas.f += datas.f2[i]
			if datas.f < datas.bestF {
				datas.x[i], datas.x[j] = datas.x[j], datas.x[i]
				backTrack(i+1, datas)
				datas.x[i], datas.x[j] = datas.x[j], datas.x[i]
			}
			datas.f -= datas.f2[i]
			datas.f1 -= datas.M[datas.x[j]][0]
		}
	}
}

type data struct {
	n     int     //作业数
	M     [][]int //M[i][j]表示第i个作业在机器j上需要处理的时间
	x     []int   //x[i]表示第i个处理的作业为x[i]
	bestX []int   //x[i]的最优值
	f1    int     //作业在机器1上完成处理的时间
	f2    []int   //f2[i]表示第i个作业在机器2上完成处理的时间
	f     int     //用于记录前i个作业在机器2上完成处理的时间和
	bestF int     //f的最优值
}
