package medium

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	var backtrack func(n_open int, n_close int, comb string)
	backtrack = func(n_open int, n_close int, comb string) {
		if n_open+n_close == 2*n {
			result = append(result, comb)
			return
		}

		if n_open < n {
			backtrack(n_open+1, n_close, comb+"(")
		}

		if n_open > n_close {
			backtrack(n_open, n_close+1, comb+")")
		}
	}

	backtrack(0, 0, "")
	return result
}
