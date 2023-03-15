package grayscaleImageCompression

import "math"

const header = 11

// Compression 灰度图像压缩，输入由n个int值组成的int切面，返回压缩后的总大小和分隔方式；段头大小为11bit，压缩后每位占据大小为(log(maxP+1))，
// 每个位取值范围为0-255
//
// 采用动态规划算法。
func Compression(bits []int) (sum int, segments []int) {
	length := len(bits)
	S := make([]int, length+1)   //记录前k个位进行压缩后的大小
	seg := make([]int, length+1) //记录第K个S分段时该分法最后一段的长度
	b := make([]int, length+1)   //记录每一位需要的位数
	S[0] = 0
	seg[0] = 0
	for i := 1; i <= length; i++ { //i代表前i个位参与压缩
		maxsize := sizes(bits[i-1])
		b[i-1] = maxsize
		S[i] = S[i-1] + maxsize //最后一段长度为1时的情况，header在最后加
		seg[i] = 1
		for j := 2; j <= i; j++ { //j代表最后一段的长度
			if maxsize < b[i-j] {
				maxsize = b[i-j]
			}
			if S[i] > S[i-j]+j*maxsize {
				S[i] = S[i-j] + j*maxsize
				seg[i] = j
			}
		}
		S[i] += header
	}
	sum = S[length]
	segments = make([]int, 0)
	for i := length; i > 0; {
		long := seg[i]
		segments = append([]int{long}, segments...)
		i -= long
	}
	return
}

func sizes(num int) int {
	result := int(math.Log2(float64(num)) + 1)
	if result < 1 {
		return 1
	}
	return result
}
