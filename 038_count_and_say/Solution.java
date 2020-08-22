class Solution {
    public String countAndSay(int n) {
        StringBuilder tmp1 = new StringBuilder("1");
        StringBuilder tmp2;
        for (int i = 2; i < n + 1; i++) {
            tmp2 = new StringBuilder();
            int count = 1;
            for (int j = 1; j < tmp1.length(); j++) {
                if (tmp1.charAt(j) == tmp1.charAt(j - 1)) count++;
                else tmp2.append(count).append(tmp1.charAt(j - 1));
            }
            tmp2.append(count).append(tmp1.charAt(tmp1.length() - 1));
            tmp1 = tmp2;
        }

        return tmp1.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.countAndSay(4));
    }
}

