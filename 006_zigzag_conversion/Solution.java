import java.util.ArrayList;

/**
 * String 为不可变对象，每次修改均需重新创建，故性能较差
 */
public class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) return s;
        
        ArrayList<StringBuilder> rows = new ArrayList<>();
        for (int i = 0; i < numRows; i++)
            rows.add(new StringBuilder());

        int curRow = 0;
        boolean downward = false;
        for (char ch : s.toCharArray()) {
            rows.get(curRow).append(ch);
            if (curRow == 0 || curRow == numRows - 1) downward = !downward;
            curRow += downward ? 1 : -1; 
        }

        StringBuilder result = new StringBuilder();
        for (StringBuilder row : rows)
            result.append(row);
        return result.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.convert("LEETCODEISHIRING", 3));
    }
}
