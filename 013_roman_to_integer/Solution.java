import java.util.HashMap;

public class Solution {
    public int getValue(char c) {
        switch (c) {
            case 'I': return 1;
            case 'V': return 5;
            case 'X': return 10;
            case 'L': return 50;
            case 'C': return 100;
            case 'D': return 500;
            case 'M': return 1000;
            default: return 0;
        }
    }

    public int romanToInt(String s) {
        int result = 0;
        int num1 = getValue(s.charAt(0));
        int num2 = 0;

        for (int i = 1; i < s.length(); i++) {
            num2 = getValue(s.charAt(i));
            if (num1 >= num2) result += num1;
            else result -= num1;

            num1 = num2;
        }
        result += num1;
        
        return result;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.romanToInt("MCMXCIV"));
    }
}
