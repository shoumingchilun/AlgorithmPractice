package addSpaces

//2109. 向字符串添加空格

func addSpaces(s string, spaces []int) string {
	spacesIndex := 0
	resultIndex := 0
	result := make([]byte, len(s)+len(spaces))
	for i := 0; i < len(s); i++ {
		if spacesIndex < len(spaces) && i == spaces[spacesIndex] {
			result[resultIndex] = []byte(" ")[0]
			resultIndex++
			spacesIndex++
		}
		result[resultIndex] = s[i]
		resultIndex++
	}
	return string(result)
}
