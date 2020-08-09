class Solution:
    def __init__(self):
        self.result = []
        self.mapping = {
            '2': ['a', 'b', 'c'],
            '3': ['d', 'e', 'f'],
            '4': ['g', 'h', 'i'],
            '5': ['j', 'k', 'l'],
            '6': ['m', 'n', 'o'],
            '7': ['p', 'q', 'r', 's'],
            '8': ['t', 'u', 'v'],
            '9': ['w', 'x', 'y', 'z']
        }

    def letterCombinations(self, digits: str) -> list:
        if not digits or len(digits) == 0: return self.result
        self.dfs(digits, '')
        return self.result

    def dfs(self, digits: str, tmp_str: str):
        if len(digits) == 0:
            self.result.append(tmp_str)
            return

        for value in self.mapping[digits[0]]:
            self.dfs(digits[1:], tmp_str + value)


if __name__ == '__main__':
    solution = Solution()
    print(solution.letterCombinations('23'))

