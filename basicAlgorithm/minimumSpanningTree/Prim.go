package minimumSpanningTree

import (
	"math"
	"workSpace/basicDataStructure"
)

type dis struct {
	source   int
	distance float64
}

func Prim(am basicDataStructure.AdjacencyMatrix) (ans basicDataStructure.AdjacencyMatrix) {
	//先获得节点数
	peers := len(am.Distances[0])
	ans = basicDataStructure.CreateAdjacencyMatrix(peers)
	//初始化一些数据，默认从0号节点出发
	distances := make([]dis, peers) //用于记录到达给点的距离
	for i := 0; i < peers; i++ {
		distances[i].distance = am.Distances[0][i]
		distances[i].source = 0
	}
	//找到要加入的节点，开始一系列操作
	for i := 1; i < peers; i++ { //运行peers-1次以加入剩余peers-1个节点
		min := math.MaxFloat64
		minPeer := 0                 //记录最近的点
		source := 0                  //通过source将最近的点加入到最小生成树
		for j := 1; j < peers; j++ { //获得最近的点
			if distances[j].distance < min && distances[j].distance != 0 { //距离为0说明已经加入
				min = distances[j].distance
				minPeer = j
				source = distances[j].source
			}
		}
		//然后加到ans
		ans.Distances[source][minPeer] = min
		ans.Distances[minPeer][source] = min
		//然后更新distances
		for j := 0; j < peers; j++ {
			if distances[j].distance > am.Distances[minPeer][j] {
				distances[j].distance = am.Distances[minPeer][j]
				distances[j].source = minPeer
			}
		}
	}
	return
}
