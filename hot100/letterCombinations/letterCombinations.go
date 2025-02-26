package letterCombinations

/*
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
17. 电话号码的字母组合
*/
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	relations := map[int][]string{
		50: {"a", "b", "c"},
		51: {"d", "e", "f"},
		52: {"g", "h", "i"},
		53: {"j", "k", "l"},
		54: {"m", "n", "o"},
		55: {"p", "q", "r", "s"},
		56: {"t", "u", "v"},
		57: {"w", "x", "y", "z"},
	}
	ret := []string{""}
	tempRet := []string{""}
	digitsSlice := []int{}
	for _, s := range digits {
		digitsSlice = append(digitsSlice, int(s))
	}
	tempRet = make([]string, len(ret), len(ret))
	copy(tempRet, ret)
	for _, key := range digitsSlice {
		end := relations[key]
		ret = []string{}
		for _, s := range end {
			for _, origin := range tempRet {
				ret = append(ret, origin+s)
			}
		}
		tempRet = make([]string, len(ret), len(ret))
		copy(tempRet, ret)
	}
	return ret
}
func letterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}
	relations := map[int][]string{
		50: {"a", "b", "c"},
		51: {"d", "e", "f"},
		52: {"g", "h", "i"},
		53: {"j", "k", "l"},
		54: {"m", "n", "o"},
		55: {"p", "q", "r", "s"},
		56: {"t", "u", "v"},
		57: {"w", "x", "y", "z"},
	}
	ret := []string{}
	digitsSlice := []int{}
	for _, s := range digits {
		digitsSlice = append(digitsSlice, int(s))
	}
	backtraceLetterCombinations2(digitsSlice, &ret, 0, "", &relations)
	return ret
}
func backtraceLetterCombinations2(digitsSlice []int, ret *[]string, index int, currentS string, relations *map[int][]string) {
	if len(digitsSlice) == index {
		*ret = append(*ret, currentS)
		return
	}
	ends := (*relations)[digitsSlice[index]]
	for _, end := range ends {
		backtraceLetterCombinations2(digitsSlice, ret, index+1, currentS+end, relations)
	}
}
