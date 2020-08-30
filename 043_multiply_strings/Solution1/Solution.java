class Solution {
    public String multiply(String num1, String num2) {
        String result = "";
        int n1 = num1.length() - 1, n2 = num2.length() - 1;
        if (num1.equals("0") || num2.equals("0")) return "0";

        int carry = 0;
        int curNum1, curNum2, tmp;

        for (int j = n2; j >= 0; j--) {
            curNum2 = num2.charAt(j) - '0';
            StringBuilder tmpRes = new StringBuilder();
            for (int i = n1; i >= 0 || carry !=0; i--) {
                curNum1 = i >= 0 ? num1.charAt(i) - '0' : 0;
                tmp = curNum1 * curNum2 + carry;
                carry = tmp / 10;
                tmpRes.append(tmp % 10);
            }
            result = addString(result, tmpRes.reverse().toString(), n2 - j);
        }

        return result;
    }

    public String addString(String s1, String s2, int digitCount) {
        StringBuilder result = new StringBuilder();
        s2 = s2 + "0".repeat(digitCount);
        int i = s1.length() - 1, j = s2.length() - 1, carry = 0;

        while (i >= 0 || j >= 0 || carry != 0) {
            int x = i >= 0 ? s1.charAt(i) - '0' : 0;
            int y = j >= 0 ? s2.charAt(j) - '0' : 0;
            int tmp = x + y + carry;
            result.append(tmp % 10);
            carry = tmp / 10;
            j--; i--;
        }

        return result.reverse().toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.multiply("123", "45"));
    }
}

