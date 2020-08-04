public class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) return false;

        int result = 0;
        int y = x;

        try {
            while (y != 0) {
                result = result * 10 + y % 10;
                y /= 10;
            }
        } catch (Exception e) {
            return false;
        }

        return x == result;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.isPalindrome(12321));
    }
}
