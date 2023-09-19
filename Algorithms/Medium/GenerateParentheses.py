class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        
        def backtrack(n_open, n_close, comb):
            if n_open + n_close == 2 * n:
                result.append(comb)
                return

            if n_open < n:
                backtrack(n_open + 1, n_close, comb + "(")

            if n_open > n_close:
                backtrack(n_open, n_close + 1, comb + ")")


        result = []
        backtrack(0, 0, "")

        return result