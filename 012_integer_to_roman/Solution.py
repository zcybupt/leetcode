class Solution:
    def __init__(self):
        self.values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        self.symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]

    def intToRoman(self, num: int) -> str:
        result = ''
        for i in range(len(self.values)):
            while num >= 0 and self.values[i] <= num:
                num -= self.values[i]
                result += self.symbols[i]

        return result


if __name__ == '__main__':
    solution = Solution()
    print(solution.intToRoman(140))
