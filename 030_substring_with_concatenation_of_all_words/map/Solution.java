import java.util.List;
import java.util.ArrayList;
import java.util.HashMap;

class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> results = new ArrayList<>();
        int strLen = s.length();
        int arrLen = words.length;
        if (strLen == 0 || arrLen == 0) return results;
        int wordLen = words[0].length();

        HashMap<String, Integer> wordMap = new HashMap<>();
        for (String word : words) {
            int value = wordMap.getOrDefault(word, 0);
            wordMap.put(word, value + 1);
        }

        HashMap<String, Integer> curMap = new HashMap<>();
        int count = 0;
        for (int i = 0; i < strLen - arrLen * wordLen + 1; i++) {
            while (count < arrLen) {
                String curWord = s.substring(i + count * wordLen, i + (count + 1) * wordLen);
                if (wordMap.containsKey(curWord)) {
                    int value = curMap.getOrDefault(curWord, 0);
                    if (value + 1 > wordMap.get(curWord)) break;
                    curMap.put(curWord, value + 1);
                } else {
                    break;
                }
                count++;
            }
            if (count == arrLen) results.add(i);
            curMap.clear();
            count = 0;
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
