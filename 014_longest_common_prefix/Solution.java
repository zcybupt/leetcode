public class Solution {
    public String longestCommonPrefix(String[] strs) {
       if (strs == null || strs.length == 0) return "";

       String commonPrefix = strs[0];
       for (int i = 1; i < strs.length; i++) {
           int j = 0;
           for (; j < commonPrefix.length() && j < strs[i].length(); j++) {
               if (commonPrefix.charAt(j) != strs[i].charAt(j)) break;
           }
           commonPrefix = commonPrefix.substring(0, j);
           if (commonPrefix.equals("")) return "";
       }

       return commonPrefix;
    }

    public static void main(String[] args) {
        String[] strs = new String[]{"flow", "flower", "flour"};
        Solution solution = new Solution();
        System.out.println(solution.longestCommonPrefix(strs));
    }
}
