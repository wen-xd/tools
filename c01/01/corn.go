package _1

func CornWay1(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	return cornWay1(arr, 0, aim)
}

func cornWay1(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}

	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += cornWay1(arr, index+1, rest-zhang*arr[index])
	}
	return ways
}

/*****************************************   method2   *************/

func CornWay2(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	size := len(arr)
	dp := make([][]int, size+1, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, aim+1, aim+1)
		for j := 0; j < aim+1; j++ {
			dp[i][j] = -1
		}
	}
	return cornWay2(arr, 0, aim, dp)
}

func cornWay2(arr []int, index, rest int, dp [][]int) int {
	if dp[index][rest] != -1 {
		return dp[index][rest]
	}
	if index == len(arr) {
		if rest == 0 {
			dp[index][rest] = 1
			return 1
		} else {
			dp[index][rest] = 0
			return 0
		}
	}

	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += cornWay1(arr, index+1, rest-zhang*arr[index])
	}
	dp[index][rest] = ways
	return ways
}

/***************** method3 ********************************************/

func CornWay3(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1, aim+1)
	}
	dp[N][0] = 1

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {

			ways := 0
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				ways += dp[index+1][rest-zhang*arr[index]]
			}
			dp[index][rest] = ways
		}
	}

	return dp[0][aim]
}

/***************** method3 ********************************************/

func CornWay4(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, aim+1, aim+1)
	}
	dp[N][0] = 1

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {

			dp[index][rest] = dp[index+1][rest]

			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index][rest-arr[index]]
			}

		}
	}

	return dp[0][aim]
}
