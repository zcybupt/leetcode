class Solution {
    public int fib(int N) {
        if (N < 2) return N;

        int a = 0, b = 1;
        while (N-- > 1) {
            int tmp = a;
            a = b;
            b += tmp;
        }

        return b;
    }

    public static void main(String[] args) {
        System.out.println(new Solution().fib(7));
    }
}
