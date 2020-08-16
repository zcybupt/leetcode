public class Solution {
    public int divide(int dividend, int divisor) {
        if (dividend == Integer.MIN_VALUE && divisor == -1) return Integer.MAX_VALUE;
        long m = Math.abs((long)dividend), n = Math.abs((long)divisor);
        int res = 0;
        int sign = (dividend < 0) ^ (divisor < 0) ? -1 : 1;
        if (n == 1) return (int) (sign == 1 ? m : -m);
        
        while (m >= n) {
            long t = n, p = 1;
            while (m >= (t << 1)) {
                t <<= 1;
                p <<= 1;
            }
            res += p;
            m -= t;
        }
        
        return sign == 1 ? res : -res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.divide(7, -3));
    }
}
