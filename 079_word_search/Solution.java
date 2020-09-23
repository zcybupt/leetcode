class Solution {
    private int m, n;
    private boolean[][] visited;
    private final int[][] directions = {{-1, 0}, {0, -1}, {1, 0}, {0, 1}};
    private char[][] board;
    private String word;
    private int wordLen;

    public boolean exist(char[][] board, String word) {
        if (board == null || board.length == 0 || board[0].length == 0) return false;
        m = board.length;
        n = board[0].length;
        this.visited = new boolean[m][n];
        this.board = board;
        this.word = word;
        this.wordLen = word.length();

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (dfs(i, j, 0)) return true;
            }
        }

        return false;
    }

    private boolean dfs(int i, int j, int strIdx) {
        if (strIdx == wordLen - 1) return board[i][j] == word.charAt(strIdx);

        if (board[i][j] != word.charAt(strIdx)) return false;

        visited[i][j] = true;
        for (int k = 0; k < 4; k++) {
            int newI = i + directions[k][0];
            int newJ = j + directions[k][1];
            if (inArea(newI, newJ) && !visited[newI][newJ]) {
                if (dfs(newI, newJ, strIdx + 1)) return true;
            }
        }
        visited[i][j] = false;

        return false;
    }

    private boolean inArea(int i, int j) {
        return 0 <= i && i < m && 0 <= j && j < n;
    }

    public static void main(String[] args) {
        char[][] board = {
            {'A', 'B', 'C', 'E'},
            {'S', 'F', 'C', 'S'},
            {'A', 'D', 'E', 'E'}
        };

        Solution solution = new Solution();
        System.out.println(solution.exist(board, "ABCCE"));
        System.out.println(solution.exist(board, "SEE"));
        System.out.println(solution.exist(board, "ABCB"));
    }
}
