class Solution:
    def __init__(self):
        self.result = []

    def generateParenthesis(self, n: int) -> list:
        self.dfs(n, n, "")
        return self.result

    def dfs(self, l: int, r: int, cur_str: str) -> None:
        if (l == 0 and r == 0):
            self.result.append(cur_str)
            return

        if (l > 0): self.dfs(l - 1, r, cur_str + '(')
        if (r > l): self.dfs(l, r - 1, cur_str + ')')


if __name__ == '__main__':
    solution = Solution()
    result = solution.generateParenthesis(3)
    for string in result:
        print(string)

