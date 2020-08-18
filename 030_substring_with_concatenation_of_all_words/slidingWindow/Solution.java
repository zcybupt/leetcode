import java.util.*;

class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> results = new ArrayList<>();
        int sLen = s.length(), arrLen = words.length;
        if (sLen == 0 || arrLen == 0) return results;
        int wordLen = words[0].length();

        Map<String, Integer> wordMap = new HashMap<>();
        for (String word : words)
            wordMap.put(word, wordMap.getOrDefault(word, 0) + 1);

        Map<String, Integer> curMap = new HashMap<>();
        int windowSize = wordLen * arrLen;
        for (int i = 0; i < wordLen; i++) {
            int start = i;
            while (start + windowSize <= sLen) {
                String str = s.substring(start, start + windowSize);
                curMap.clear();
                int j = arrLen;
                while (j > 0) {
                    String curWord = str.substring((j - 1) * wordLen, j * wordLen);
                    int count = curMap.getOrDefault(curWord, 0) + 1;
                    if (count > wordMap.getOrDefault(curWord, 0)) break;
                    curMap.put(curWord, count);
                    j--;
                }
                if (j == 0) results.add(start);
                start += wordLen * Math.max(j, 1);
            }
        }

        return results;
    }

    public static void main(String[] args) {
        String s = "barfoothefoobarman";
        String[] words = new String[]{"foo", "bar"};
        Solution solution = new Solution();
        List<Integer> results = solution.findSubstring(s, words);
        for (int result : results) {
            System.out.print(result + " ");
        }
    }
}
