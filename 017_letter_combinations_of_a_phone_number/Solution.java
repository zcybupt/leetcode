import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    List<String> result = new ArrayList<>();
    HashMap<String, String[]> mapping = new HashMap<>() {
        {
            put("2", new String[]{"a", "b", "c"});
            put("3", new String[]{"d", "e", "f"});
            put("4", new String[]{"g", "h", "i"});
            put("5", new String[]{"j", "k", "l"});
            put("6", new String[]{"m", "n", "o"});
            put("7", new String[]{"p", "q", "r", "s"});
            put("8", new String[]{"t", "u", "v"});
            put("9", new String[]{"w", "x", "y", "z"});
        }
    };

    public List<String> letterCombinations1(String digits) {
        if (digits == null || digits.length() == 0) return null;

        for (String digit : digits.split("")) {
            if (result.size() < 3) {
                result.addAll(Arrays.asList(mapping.get(digit)));
                continue;
            }

            List<String> tmpList = new ArrayList<>(result);
            for (String ele : tmpList) {
                for (String newEle : mapping.get(digit)) {
                    result.add(ele + newEle);
                }
                result.remove(ele);
            }
        }
        
        return result;
    }

    public List<String> letterCombinations2(String digits) {
        if (digits == null || digits.length() == 0) return result;
        dfs(digits, "");
        return result;
    }

    public void dfs(String digits, String str) {
        if (digits.length() == 0) {
            result.add(str);
            return;
        }

        for (String value : mapping.get(digits.substring(0, 1))) {
            dfs(digits.substring(1), str + value);
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        List<String> result = solution.letterCombinations2("234");
        for (String ele : result) {
            System.out.print(ele + " ");
        }
    }
}
