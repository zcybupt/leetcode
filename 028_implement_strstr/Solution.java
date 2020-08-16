public class Solution {
    public int[] getNextTable(String needle) {
        int strLen = needle.length();
        int[] next = new int[strLen];
        next[0] = -1;
        int lp = -1, rp = 0;

        while (rp < strLen - 1) {
            if (lp == -1 || needle.charAt(lp) == needle.charAt(rp)) {
                next[++rp] = ++lp;
            } else {
                lp = next[lp];
            }
        }

        return next;
    }

    public int strStr(String haystack, String needle) {
        int haystackLen = haystack.length();
        int needleLen = needle.length();

        if (needleLen == 0) return 0;
        int haystackCur = 0, needleCur = 0;
        int[] next = getNextTable(needle);

        while (haystackCur < haystackLen && needleCur < needleLen) {
            if (haystack.charAt(haystackCur) == needle.charAt(needleCur)) {
                haystackCur++; needleCur++;
            } else {
                if (needleCur == 0) haystackCur++;
                else needleCur = next[needleCur];
            }
        }

        return needleCur == needleLen ? haystackCur - needleCur : -1;
    }

    public static void main(String[] args) {
        String haystack = "hello";
        String needle = "ll";
        Solution solution = new Solution();
        System.out.println(solution.strStr(haystack, needle));
    }
}
