class Solution {
    final int[] facts = {1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880};

    public int getFactorial(int n) {
        if (n < 10) return facts[n];
        return n * getFactorial(n - 1);
    }

    public String getPermutation(int n, int k) {
        StringBuilder result = new StringBuilder();
        boolean[] isVisited = new boolean[n];
        k--;
        for (int i = 0; i < n; i++) {
            int tmpFact = getFactorial(n - 1 - i);
            int cnt = k / tmpFact;
            k %= tmpFact;
            for (int j = 0; j < n; j++) {
                if (isVisited[j]) continue;
                if (cnt-- == 0) {
                    isVisited[j] = true;
                    result.append(j + 1);
                    break;
                }
            }
        }

        return String.valueOf(result);
    }

    public static void main(String[] args) {
        System.out.println(new Solution().getPermutation(4, 9));
    }
}
