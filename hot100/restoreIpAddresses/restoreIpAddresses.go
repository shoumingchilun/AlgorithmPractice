package restoreIpAddresses

import (
	"fmt"
	"strconv"
)

/*
https://leetcode.cn/problems/restore-ip-addresses/description/
93. 复原 IP 地址
*/
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}
	ret := []string{}
	backtrace([]int{}, 0, s, &ret)
	return ret
}

func backtrace(currentNums []int, index int, s string, ret *[]string) {
	if len(currentNums) == 4 {
		if index == len(s) {
			*ret = append(*ret, fmt.Sprintf("%d.%d.%d.%d", currentNums[0], currentNums[1], currentNums[2], currentNums[3]))
		}
		return
	}
	if s[index] == 48 {
		temp := append([]int{}, append(currentNums, 0)...)
		backtrace(temp, index+1, s, ret)
	} else {
		unusedCount := len(s) - index //尚未使用的字符串长度
		if unusedCount > 3*(4-len(currentNums)) || unusedCount < (4-len(currentNums)) {
			//太多或者不够分的情况
			return
		}
		if unusedCount == (4 - len(currentNums)) {
			//刚好够分的情况
			temp := append([]int{}, append(currentNums, int(s[index])-48)...)
			backtrace(temp, index+1, s, ret)
		} else if unusedCount == (4 - len(currentNums) + 1) {
			//能分1~2位
			//分一位
			temp1 := append([]int{}, append(currentNums, int(s[index])-48)...)
			backtrace(temp1, index+1, s, ret)
			//分两位
			str := s[index : index+2]
			num, _ := strconv.Atoi(str)
			temp2 := append([]int{}, append(currentNums, num)...)
			backtrace(temp2, index+2, s, ret)
		} else {
			//能分1~3位
			//分一位
			temp1 := append([]int{}, append(currentNums, int(s[index])-48)...)
			backtrace(temp1, index+1, s, ret)
			//分两位
			str := s[index : index+2]
			num, _ := strconv.Atoi(str)
			temp2 := append([]int{}, append(currentNums, num)...)
			backtrace(temp2, index+2, s, ret)
			//分三位
			str = s[index : index+3]
			num, _ = strconv.Atoi(str)
			if num <= 255 {
				temp3 := append([]int{}, append(currentNums, num)...)
				backtrace(temp3, index+3, s, ret)
			}
		}
	}
}
