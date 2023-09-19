package _1

func MaxValue(w, v []int, bag int) int {
	return process(w, v, 0, bag)
}

func process(w, v []int, index, rest int) int {
	if rest < 0 { // 3. base case 1  要了价值后bag是否承受不了
		return -1
	}
	if index == len(w) { // 没货物了
		return 0
	}
	p1 := process(w, v, index+1, rest) // 1. 当前货物没要，后面的价值
	p2 := -1
	p2Next := process(w, v, index+1, rest-w[index]) // 2. 要了，求后面的价值
	if p2Next != -1 {
		p2 = v[index] + p2Next // p2以及后面的价值和
	}
	return max(p1, p2) // 返回最大价值
}

/*********************************  method2   *****************************************/

func DpWays(w, v []int, bag int) int {
	size := len(w)
	dp := make([][]int, size+1, size+1)
	dp[size] = make([]int, bag+1, bag+1)

	for index := size - 1; index >= 0; index-- {
		dp[index] = make([]int, bag+1, bag+1)
		for rest := 0; rest <= bag; rest++ {

			p1 := dp[index+1][rest]
			p2 := -1
			if rest-w[index] >= 0 {
				p2 = dp[index+1][rest-w[index]] + v[index]
			}
			dp[index][rest] = max(p1, p2)
		}
	}

	return dp[0][bag]
}
