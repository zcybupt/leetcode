class Solution {
    public String getPermutation(int n, int k) {
        if (n == 0) return "";
        char[] permutation = new char[n];
        for (int i = 0; i < n; i++) permutation[i] = (char) (i + '1');

        for (int i = 1; i < k; i++) nextPermutation(permutation);

        return String.valueOf(permutation);
    }

    public boolean nextPermutation(char[] permutation) {
        int arrLen = permutation.length;
        int i = arrLen - 2, j = arrLen - 1, k = arrLen - 1;

        while (i >= 0 && permutation[i] >= permutation[j]) {
            i--; j--;
        }

        if (i == -1) return false;

        while (permutation[i] >= permutation[k]) k--;
        swap(permutation, i, k);

        for (i = j, j = arrLen - 1; i < j; i++, j--)
            swap(permutation, i, j);

        return true;
    }

    public void swap(char[] permutation, int i, int j) {
        char tmp = permutation[i];
        permutation[i] = permutation[j];
        permutation[j] = tmp;
    }

    public static void main(String[] args) {
        System.out.println(new Solution().getPermutation(4, 9));
    }
}
