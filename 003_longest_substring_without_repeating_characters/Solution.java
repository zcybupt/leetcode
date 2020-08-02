import java.util.HashSet;

public class Solution {
    public int lengthOfLongestSubstring(String s) {
        HashSet<Character> charSet = new HashSet<>();
        int length = s.length();
        int rk = -1, ans = 0;

        for (int i = 0; i < length; i++) {
            if (i != 0) {
                charSet.remove(s.charAt(i - 1));
            }
            while (rk + 1 < length && !charSet.contains(s.charAt(rk + 1))) {
                charSet.add(s.charAt(rk + 1));
                rk++;
            }
            ans = Math.max(ans, rk - i + 1);
        }

        return ans;
    }

    public static void main(String[] args) {
        String s = "abcabcbb";
        Solution solution = new Solution();
        System.out.println(solution.lengthOfLongestSubstring(s));
    }
}
