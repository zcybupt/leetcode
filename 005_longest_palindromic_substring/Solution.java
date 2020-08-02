public class Solution {
    public String preprocess(String s) {
        String result = "^";
        for (int i = 0; i < s.length(); i++) {
            result += "#" + s.charAt(i);
        }

        return result + "#$";
    }

    public String longestPalindrome(String s) {
        if (s.length() == 0) return "";

        String tmpStr = preprocess(s);
        int len = tmpStr.length();
        int[] p = new int[len];
        int C = 0, R = 0;
        int maxLen = -1;
        int index = 0;

        for (int i = 1; i < len - 1; i++) {
            p[i] = R > i ? Math.min(p[2 * C - i], R - i) : 1;

            while (tmpStr.charAt(i + p[i]) == tmpStr.charAt(i - p[i]))
                p[i]++;

            if (R < i + p[i]) {
                R = i + p[i];
                C = i;
            }

            if (maxLen < p[i] - 1) {
                maxLen = p[i] - 1;
                index = i;
            }
        }

        int start = (index - maxLen) / 2;
        return s.substring(start, start + maxLen);
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.longestPalindrome("babad"));
    }
}
