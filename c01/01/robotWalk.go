package _1

func Way1(N, M, K, P int) int {
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}
	return walk1(N, M, K, P)
}

// N 一共有几个位置 固定参数 1-N
// cur 当前的位置
// rest 还剩几步可以走
// p 目标位置
func walk1(N, cur, rest, P int) int {
	if rest == 0 {
		if P == cur {
			return 1
		}
		return 0
	}

	if cur == 1 {
		return walk1(N, 2, rest-1, P)
	}
	if cur == N {
		return walk1(N, N-1, rest-1, P)
	}
	return walk1(N, cur-1, rest-1, P) + walk1(N, cur+1, rest-1, P)
}

/************************  method2  *********************************/

func Way2(N, M, K, P int) int {
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}
	dp := make([][]int, N+1, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, K+1, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = -1
		}
	}

	return walk2(N, M, K, P, dp)
}

// N 一共有几个位置 固定参数 1-N
// cur 当前的位置
// rest 还剩几步可以走
// p 目标位置
func walk2(N, cur, rest, P int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	if rest == 0 {
		if P == cur {
			dp[cur][rest] = 1
			return 1
		}
		dp[cur][rest] = 0
		return 0
	}

	if cur == 1 {
		dp[cur][rest] = walk2(N, 2, rest-1, P, dp)
		return dp[cur][rest]
	}
	if cur == N {
		dp[cur][rest] = walk2(N, N-1, rest-1, P, dp)
		return dp[cur][rest]
	}
	dp[cur][rest] = walk2(N, cur-1, rest-1, P, dp) + walk2(N, cur+1, rest-1, P, dp)
	return dp[cur][rest]
}

/************************  method3  *********************************/

func RobotDpWays(N, cur, rest, P int) int {
	if N < 2 || rest < 1 || cur < 1 || cur > N || P < 1 || P > N {
		return 0
	}

	dp := make([][]int, N+1, N+1)
	for index := 1; index <= N; index++ {
		dp[index] = make([]int, rest+1, rest+1)
	}
	dp[P][0] = 1

	for index := 1; index <= rest; index++ { // 列
		for j := 1; j <= N; j++ { // 行

			if j == 1 {
				dp[j][index] = dp[j+1][index-1]
				continue
			}
			if j == N {
				dp[j][index] = dp[j-1][index-1]
				continue
			}
			dp[j][index] = dp[j+1][index-1] + dp[j-1][index-1]
		}
	}

	return dp[cur][rest]
}
