class Solution {
    public double quickMul(double x, long n) {
        double result = 1.0;
        while (n >= 1) {
            if (n % 2 == 1) result *= x;
            x *= x;
            n /= 2;
        }

        return result;
    }

    public double myPow(double x, int n) {
        return n >= 0 ? quickMul(x, n) : 1.0 / quickMul(x, -(long) n);
    }

    public static void main(String[] args) {
        System.out.println(new Solution().myPow(2, -2147483648));
    }
}
