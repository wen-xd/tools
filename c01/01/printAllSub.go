package _1

func Subs(str string) []string {
	var ans []string
	path := ""
	process1(str, 0, &ans, path)
	return ans
}

func process1(str string, index int, ans *[]string, path string) {
	if index == len(str) {
		*ans = append(*ans, path)
		return
	}
	no := path
	process1(str, index+1, ans, no)

	yes := path + string(str[index])
	process1(str, index+1, ans, yes)
}

func Permutation(str string) []string {
	if str == "" {
		return nil
	}
	var ans []string
	process2(&str, 0, &ans)
	return ans
}
func process2(str *string, i int, ans *[]string) {
	if i == len(*str) {
		*ans = append(*ans, *str)
	}

	var visited [26]bool
	for j := i; j < len(*str); j++ {
		if !visited[(*str)[j]-'a'] {
			visited[(*str)[j]-'a'] = true
			swap(str, i, j)
			process2(str, i+1, ans)
			swap(str, i, j)
		}
	}
}

func swap(str *string, i, j int) {
	if i == j {
		return
	}
	a := []byte(*str)
	a[i], a[j] = (*str)[j], (*str)[i]
	*str = string(a)
}
