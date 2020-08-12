import java.util.List;
import java.util.ArrayList;

public class Solution {
    List<String> result = new ArrayList<>();

    public List<String> generateParenthesis(int n) {
       dfs(n, n, "");

       return result;
    }

    public void dfs(int l, int r, String curStr) {
        if (l == 0 && r == 0) {
            result.add(curStr);
            return;
        }

        if (l > 0) dfs(l - 1, r, curStr + "(");
        if (r > l) dfs(l, r - 1, curStr + ")");
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        List<String> res = solution.generateParenthesis(3);
        for (String str : res) {
            System.out.println(str);
        }
    }
}
