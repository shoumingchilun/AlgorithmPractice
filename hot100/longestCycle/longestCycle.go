package longestCycle

func longestCycle(edges []int) int {
	//记录每个节点的来源节点
	inEdges := make(map[int][]int)
	for start, end := range edges {
		if end == -1 {
			continue
		}
		if starts, exist := inEdges[end]; exist {
			inEdges[end] = append(starts, start)
		} else {
			inEdges[end] = []int{start}
		}
	}
	//用于记录过滤了无入度或无出度的节点
	filteredNodes := make(map[int]struct{})
	zeroOut := make([]int, 0)
	//记录无出度的节点
	for i := 0; i < len(edges); i++ {
		filteredNodes[i] = struct{}{}
		if edges[i] == -1 {
			zeroOut = append(zeroOut, i)
		}
	}
	//过滤无出度的节点
	for len(zeroOut) > 0 {
		tempZeroOut := make([]int, 0)
		for _, i := range zeroOut {
			delete(filteredNodes, i)
			tempZeroOut = append(tempZeroOut, inEdges[i]...)
		}
		zeroOut = tempZeroOut
	}
	//至此，已过滤无出度，接下来过滤无入度(拓扑排序)
	//记录每个节点的入度数量
	inEdgeNums := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		inEdgeNums[i] = 0
	}
	for index, starts := range inEdges {
		inEdgeNums[index] = len(starts)
	}
	zeroIn := make([]int, 0)
	for index, num := range inEdgeNums {
		if num == 0 {
			zeroIn = append(zeroIn, index)
		}
	}
	for len(zeroIn) > 0 {
		tempZeroIn := make([]int, 0)
		for _, node := range zeroIn {
			delete(filteredNodes, node)
			inEdgeNums[edges[node]]--
			if inEdgeNums[edges[node]] == 0 {
				tempZeroIn = append(tempZeroIn, edges[node])
			}
		}
		zeroIn = tempZeroIn
	}
	//fmt.Println(filteredNodes)
	if len(filteredNodes) == 0 {
		return -1
	}
	//使用并查集，处理全部过滤后节点
	fa := make([]int, len(edges))
	for index, _ := range fa {
		fa[index] = index
	}

	for node, _ := range filteredNodes {
		union(node, edges[node], fa)
	}
	resultCountMap := make(map[int]int)
	maxCount := 0
	for node, _ := range fa {
		fa := find(node, fa)
		resultCountMap[fa]++
		if resultCountMap[fa] > maxCount {
			maxCount = resultCountMap[fa]
		}
	}

	return maxCount
}

func find(x int, fa []int) int {
	if fa[x] != x {
		fa[x] = find(fa[x], fa) // 路径压缩
	}
	return fa[x]
}
func union(i int, j int, fa []int) {
	iFather := find(i, fa) //找到i的祖先节点
	jFather := find(j, fa) //找到j的祖先节点
	fa[iFather] = jFather  //i的祖先指向j的祖先
}
