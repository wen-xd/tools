package _1

func Queen(n int) int {
	if n < 1 {
		return 0
	}
	var record = make([]int, n, n)
	return m1(0, n, record)
}

func m1(i, n int, record []int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) { // 验证在i行j列 是否可以安放皇后
			record[i] = j             // 记录每一行皇后所在的位置
			res += m1(i+1, n, record) // 在下一行继续递归
		}
	}
	return res
}

/*
*
i : 之前的行
j : 列
record : 保存之前每一行所在的列
*/
func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ { // 遍历之前的每一行
		if j == record[k] || abs(record[k]-j) == abs(i-k) { // 判断j列是否在之前已经存在或者同斜线
			return false
		}
	}
	return true
}

func abs[T int | float64](x T) T {
	if x > 0 {
		return x
	}
	return -x
}

/******************************   method2  ****************************************/

func Queen2(n int) int {
	if n < 1 || n > 32 { // 32是极限
		return 0
	}
	limit := n
	if limit == 32 {
		limit = -1 // 相当于32位 1...11
	} else {
		limit = (1 << n) - 1 // 例如：4皇后时要得到4个1 0...1111，这时候需要左移再减1
	}
	return m2(limit, 0, 0, 0)
}

// param：总限制，列，   左，  右限制
func m2(limit, colLim, lLim, rlim int) int {
	if limit == colLim {
		return 1
	}
	pos := limit & (^(colLim | lLim | rlim)) // 1表示可以放皇后
	rightOne := 0
	res := 0
	for pos != 0 { // for 用来依次取出pos最右侧的 1
		rightOne = pos & -pos // 每次取出右侧的第一个 1
		pos ^= rightOne

		res += m2(limit,
			colLim|rightOne,
			(lLim|rightOne)<<1,
			(rlim|rightOne)>>1)
	}
	return res
}
