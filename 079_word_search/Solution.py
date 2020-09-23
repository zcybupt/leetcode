class Solution:
    def exist(self, board: list, word: str) -> bool:
        m, n = len(board), len(board[0])
        word_len = len(word)
        directions = [[-1, 0], [0, -1], [1, 0], [0, 1]]
        visited = [[False] * n for _ in range(m)]

        def dfs(i: int, j: int, str_idx: int) -> bool:
            if str_idx == word_len - 1: return board[i][j] == word[str_idx]

            if board[i][j] != word[str_idx]: return False

            visited[i][j] = True
            for direction in directions:
                new_i = i + direction[0]
                new_j = j + direction[1]
                if in_area(new_i, new_j) and not visited[new_i][new_j]:
                    if dfs(new_i, new_j, str_idx + 1):
                        return True
            visited[i][j] = False

            return False

        def in_area(i: int, j: int):
            return 0 <= i < m and 0 <= j < n

        for i in range(m):
            for j in range(n):
                if dfs(i, j, 0): return True

        return False


if __name__ == '__main__':
    board = [
        ['A','B','C','E'],
        ['S','F','C','S'],
        ['A','D','E','E']
    ]

    solution = Solution()
    print(solution.exist(board, 'ABCCED'))
    print(solution.exist(board, 'SEE'))
    print(solution.exist(board, 'ABCB'))

