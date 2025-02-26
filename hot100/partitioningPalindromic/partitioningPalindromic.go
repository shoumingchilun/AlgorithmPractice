package partitioningPalindromic

/*
https://leetcode.cn/problems/palindrome-partitioning/description/
131. 分割回文串
*/

func partitioningPalindromic(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	ret := [][]string{}
	backtrace([]string{}, 0, s, &ret)
	return ret
}

func backtrace(currentStrings []string, index int, s string, ret *[][]string) {
	if index == len(s) {
		*ret = append(*ret, currentStrings)
	}
	for i := 1; index+i <= len(s); i++ {
		if judgePalindromic(s[index : index+i]) {
			temp := append([]string{}, append(currentStrings, s[index:index+i])...)
			backtrace(temp, index+i, s, ret)
		}
	}
}
func judgePalindromic(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
