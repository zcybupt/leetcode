class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1: return s

        rows = [''] * numRows
        cur_row = 0
        downward = False
        for ch in s:
            rows[cur_row] += ch
            if cur_row in [0, numRows - 1]: downward = not downward
            cur_row += 1 if downward else -1

        return ''.join(rows)

    def convert2(self, s: str, numRows: int) -> str:
        if numRows == 1: return s

        rows = [[] for i in range(numRows)]
        cur_row = 0
        downward = False
        for ch in s:
            rows[cur_row].append(ch)
            if cur_row in [0, numRows - 1]: downward = not downward
            cur_row += 1 if downward else -1
        return ''.join([''.join(x) for x in rows])


if __name__ == '__main__':
    solution = Solution()
    print(solution.convert2('LEETCODEISHIRING', 3))
