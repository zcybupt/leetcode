class Solution:
    def countAndSay(self, n: int) -> str:
        tmp1 = '1'
        for i in range(2, n + 1):
            tmp2 = ''
            count = 1
            for j in range(1, len(tmp1)):
                if tmp1[j] == tmp1[j - 1]: count += 1
                else:
                    tmp2 += str(count) + tmp1[j - 1]
                    count = 1
            tmp2 += str(count) + tmp1[-1]
            tmp1 = tmp2

        return tmp1


if __name__ == '__main__':
    solution = Solution()
    print(solution.countAndSay(4))

