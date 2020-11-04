class Solution {
    public String reverseWords(String s) {
        s = s.trim();
        int sLen = s.length();
        if (sLen <= 1) return s;
        StringBuilder result = new StringBuilder();
        int i = sLen - 1, j = i;
        while (i >= 0) {
            while (i >= 0 && s.charAt(i) != ' ') i--;
            result.append(s.substring(i + 1, j + 1) + " ");
            while (i >= 0 && s.charAt(i) == ' ') i--;
            j = i;
        }
        return result.toString().trim();
    }

    public static void main(String[] args) {
        String input = "Stay hungry, stay foolish.";
        System.out.println(new Solution().reverseWords(input));
    }
}
