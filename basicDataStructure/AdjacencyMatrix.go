package basicDataStructure

import "workSpace/convenientMethod"

// AdjacencyMatrix 点数量为n的邻接矩阵
type AdjacencyMatrix struct {
	Distances [][]float64
}

// CreateAdjacencyMatrix 创建一个点数量为n且各点距离为0的邻接矩阵
func CreateAdjacencyMatrix(n int) AdjacencyMatrix {
	var am = AdjacencyMatrix{convenientMethod.Create2DSlice[float64](n)}
	return am
}
