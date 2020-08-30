class Solution {
    public String multiply(String num1, String num2) {
        int n1 = num1.length(), n2 = num2.length();
        if (num1.equals("0") || num2.equals("0")) return "0";
        int[] result = new int[n1 + n2];
        int curNum1, curNum2, tmp;

        for (int j = n2 - 1; j >= 0; j--) {
            curNum2 = num2.charAt(j) - '0';
            for (int i = n1 - 1; i >= 0; i--) {
                curNum1 = num1.charAt(i) - '0';
                System.out.println("i: " + i + " j: " + j);
                tmp = curNum1 * curNum2 + result[i + j + 1];
                result[i + j + 1] = tmp % 10;
                result[i + j] += tmp / 10;
            }
        }

        StringBuilder sb = new StringBuilder();
        for (int k = 0; k < result.length; k++) {
            if (k == 0 && result[k] == 0) continue;
            sb.append(result[k]);
        }

        return sb.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.multiply("123", "45"));
    }
}

