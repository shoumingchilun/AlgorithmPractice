package EasyAlgorithmProblem

//题目：当前有面值分别为 2 角 5 分，1 角，5 分，1 分的硬币，请给出找 n 分钱的最佳方案（要求找出的硬币数目最少）。

func FindRetail(n int) (retail [4]int) {
	retail[0] = n / 25
	n = n % 25
	retail[1] = n / 10
	n = n % 10
	retail[2] = n / 5
	retail[3] = n % 5
	return
}
