class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        result = ''
        if num1 == '0' or num2 == '0': return '0'
        n1, n2 = len(num1) - 1, len(num2) - 1
        carry = 0

        for j in range(n2, -1, -1):
            cur_num2 = int(num2[j])
            tmp_result = []
            i = n1
            while i >= 0 or carry != 0:
                cur_num1 = int(num1[i]) if i >= 0 else 0
                tmp = cur_num1 * cur_num2 + carry
                carry = tmp // 10
                tmp_result.append(str(tmp % 10))
                i -= 1
            result = self.add_strings(result, ''.join(reversed(tmp_result)), n2 - j)

        return result

    def add_strings(self, s1: str, s2: str, digit_count: int) -> str:
        s2 = s2 + ''.join(['0' for i in range(digit_count)])
        n1, n2 = len(s1) - 1, len(s2) - 1
        carry = 0
        result = []

        while n1 >= 0 or n2 >= 0 or carry != 0:
            x = int(s1[n1]) if n1 >= 0 else 0
            y = int(s2[n2]) if n2 >= 0 else 0
            tmp = x + y + carry
            carry = tmp // 10
            result.append(str(tmp % 10))
            n1 -= 1
            n2 -= 1

        return ''.join(reversed(result))


if __name__ == '__main__':
    solution = Solution()
    print(solution.multiply('123', '45'))

