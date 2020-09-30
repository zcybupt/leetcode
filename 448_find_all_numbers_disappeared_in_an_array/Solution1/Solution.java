import java.util.*;

class Solution {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        List<Integer> result = new ArrayList<>();
        if (nums == null || nums.length == 0) return result;
        int[] tmpArr = new int[nums.length];

        for (int num : nums) tmpArr[num - 1]++;
        for (int i = 0; i < tmpArr.length; i++) {
            if (tmpArr[i] == 0) result.add(i + 1);
        }

        return result;
    }

    public static void main(String[] args) {
        int[] nums = {4, 3, 2, 7, 8, 2, 3, 1};
        System.out.println(new Solution().findDisappearedNumbers(nums));
    }
}
