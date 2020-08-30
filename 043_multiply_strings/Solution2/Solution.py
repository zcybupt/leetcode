class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == '0' or num2 == '0': return '0'
        n1, n2 = len(num1), len(num2)
        result = [0] * (n1 + n2)

        for i in range(n1 - 1, -1, -1):
            x = int(num1[i])
            for j in range(n2 - 1, -1, -1):
                y = int(num2[j])
                tmp = x * y + result[i + j + 1]
                result[i + j + 1] = tmp % 10
                result[i + j] += tmp // 10

        if result[0] == 0: result = result[1:]
        return ''.join([str(i) for i in result])


if __name__ == '__main__':
    solution = Solution()
    print(solution.multiply("123", "45"))
