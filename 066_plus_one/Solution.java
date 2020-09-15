class Solution {
    public int[] plusOne(int[] digits) {
        int arrLen = digits.length;
        for (int i = arrLen - 1; i >= 0; i--) {;
            digits[i] = (digits[i] + 1) % 10;
            if (digits[i] != 0) return digits;
        }

        int[] result = new int[arrLen + 1];
        result[0] = 1;
        return result;
    }

    public int[] plusOne2(int[] digits) {
        for (int i = digits.length - 1; i >= 0; i--) {
            if (digits[i] != 9) {
                digits[i]++;
                return digits;
            }
            digits[i] = 0;
        }

        int[] result = new int[digits.length + 1];
        result[0] = 1;
        return result;
    }

    public static void main(String[] args) {
        int[] digits = {9, 9, 8, 9};
        int[] result = new Solution().plusOne(digits);
        for (int digit : result)
            System.out.print(digit + " ");
    }
}
