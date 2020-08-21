class Solution:
    def isValidSudoku(self, board: list) -> bool:
        rows = [[0 for i in range(9)] for i in range(9)]
        columns = [[0 for i in range(9)] for i in range(9)]
        boxes = [[0 for i in range(9)] for i in range(9)]

        for i in range(9):
            for j in range(9):
                num = board[i][j]
                if num == '.': continue
                num = int(num) - 1
                box_index = (i // 3) * 3 + j // 3

                rows[i][num] += 1
                columns[j][num] += 1
                boxes[box_index][num] += 1

                if rows[i][num] > 1 or columns[j][num] > 1 or boxes[box_index][num] > 1:
                    return False

        return True


if __name__ == '__main__':
    board = [
        ["5","3",".",".","7",".",".",".","."],
        ["6",".",".","1","9","5",".",".","."],
        [".","9","8",".",".",".",".","6","."],
        ["8",".",".",".","6",".",".",".","3"],
        ["4",".",".","8",".","3",".",".","1"],
        ["7",".",".",".","2",".",".",".","6"],
        [".","6",".",".",".",".","2","8","."],
        [".",".",".","4","1","9",".",".","5"],
        [".",".",".",".","8",".",".","7","9"]
    ]
    solution = Solution()
    print(solution.isValidSudoku(board))

