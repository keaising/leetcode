package parenthes

func generateParenthesis(n int) []string {
	ret := make([]string, 0, 0)
	dfs(0, 0, "", n, &ret)
	return ret
}

func dfs(left, right int, s string, n int, ret *[]string) {
	if right == n {
		*ret = append(*ret, s)
	} else {
		if left < n {
			dfs(left+1, right, s+"(", n, ret)
		}
		if right < left {
			dfs(left, right+1, s+")", n, ret)
		}
	}
}
