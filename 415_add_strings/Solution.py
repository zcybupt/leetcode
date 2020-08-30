class Solution:
    def addString(self, num1: str, num2: str) -> str:
        i, j = len(num1) - 1, len(num2) - 1
        result = []
        carry = 0

        while i >= 0 or j >= 0 or carry != 0:
            x = int(num1[i]) if i >= 0 else 0
            y = int(num2[j]) if j >= 0 else 0
            tmp = x + y + carry
            result.append(tmp % 10)
            carry = tmp // 10
            i -= 1
            j -= 1

        return ''.join([str(r) for r in reversed(result)])


if __name__ == '__main__':
    solution = Solution()
    print(solution.addString('123', '98'))

