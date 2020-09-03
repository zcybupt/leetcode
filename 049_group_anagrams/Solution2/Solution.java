import java.util.*;

class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        int[] primes = new int[]{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101};
        Map<Integer, ArrayList<String>> map = new HashMap<>();

        for (String str : strs) {
            int key = 1;
            for (char ch : str.toCharArray()) key *= primes[ch - 'a'];
            if (!map.containsKey(key)) map.put(key, new ArrayList<>());
            map.get(key).add(str);
        }

        return new ArrayList<>(map.values());
    }

    public static void main(String[] args) {
        String[] strs = new String[]{"eat", "tea", "tan", "ate", "nat", "bat"};
        List<List<String>> results = new Solution().groupAnagrams(strs);
        for (List<String> result : results) {
            for (String tmp : result) {
                System.out.print(tmp + " ");
            }
            System.out.println();
        }
    }
}
