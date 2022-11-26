package singleSourceShortestPaths

import (
	"math"
	"workSpace/basicDataStructure"
)

// Dijkstra 传入一个邻接矩阵，传入一个指定的出发点，传出这个出发点到其他点的最短路径
func Dijkstra(am basicDataStructure.AdjacencyMatrix, p int) (sd []float64) {
	numbers := len(am.Distances[0])
	sd = make([]float64, numbers)
	ps := make([]int, numbers) //ps存储已到达的点集
	ps[0] = p
	sd[p] = 0
	for j := 0; j < numbers; j++ {
		sd[j] = am.Distances[p][j]
	}
	for i := 1; i < numbers; i++ { //i用来计数到达点集中点的数量
		index := sp(sd, ps) //决定加入的点
		ps[i] = index       //然后更新最短距离
		for j := 0; j < numbers; j++ {
			nd := am.Distances[ps[i]][j] + sd[ps[i]]
			if sd[j] > nd {
				sd[j] = nd
			}
		}
		//fmt.Println(sd)
	}
	return
}

// sp 获得float中的除了ps之外最小的值
func sp(sd []float64, ps []int) int {
	sd2 := make([]float64, len(sd))
	for i := 0; i < len(sd); i++ {
		sd2[i] = sd[i]
	}
	for i := 0; i < len(ps); i++ {
		sd2[ps[i]] = math.MaxFloat64
	}
	min := math.MaxFloat64
	index := 0
	for i := 0; i < len(sd); i++ {
		if min > sd2[i] {
			min = sd2[i]
			index = i
		}
	}
	//fmt.Println(index)
	return index
}
