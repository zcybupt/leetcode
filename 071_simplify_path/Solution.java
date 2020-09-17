class Solution {
    public String simplifyPath(String path) {
        if (path == null || path.length() == 0) return "";

        String[] pathArr = path.split("/");
        String[] stack = new String[pathArr.length];
        int index = 0;

        for (String tmp : pathArr) {
            if (tmp.equals("") || tmp.equals("."))
                continue;
            if (tmp.equals("..")) {
                index = index > 0 ? index - 1 : 0;
                continue;
            }
            stack[index++] = tmp;
        }

        StringBuilder result = new StringBuilder();
        for (int i = 0; i < index; i++) {
            if (stack[i] == null) continue;
            result.append("/").append(stack[i]);
        }

        return index == 0 ? "/" : result.toString();
    }

    public static void main(String[] args) {
        String[] testCases = {
            "/...",
            "/home/",
            "/../",
            "/home//foo/",
            "/a/../../b/../c//.//",
            "/a//b////c/d//././/..",
            "/a/./b/../../c/",
        };

        Solution solution = new Solution();
        for (String testCase : testCases) {
            System.out.println(solution.simplifyPath(testCase));
        }
    }
}
