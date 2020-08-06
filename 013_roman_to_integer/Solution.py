class Solution:
    def romanToInt(self, s: str) -> int:
        mapping = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

        result = 0
        num1 = mapping[s[0]]
        for i, v in enumerate(s[1:]):
            num2 = mapping[v]
            if num1 >= num2: result += num1
            else: result -= num1
            num1 = num2
        result += num1

        return result


if __name__ == '__main__':
    solution = Solution()
    print(solution.romanToInt('MCMXCIV'))

