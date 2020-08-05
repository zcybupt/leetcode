import java.util.HashMap;

public class Solution {
    private static int[] values = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
    private static String[] symbols = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

    public String intToRoman(int num) {
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < values.length && num >= 0; i++) {
            while (values[i] <= num) {
                num -= values[i];
                String str = symbols[i];
                result.append(str);
            }
        }

        return result.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.intToRoman(140));
    }
}
