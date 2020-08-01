import java.util.HashMap;

public class Solution {
    public int[] twoSum(int[] nums, int target) {
        if (nums == null || nums.length == 0) {
            return null;
        }

        int[] indices = new int[2];
        HashMap<Integer, Integer> hashMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (hashMap.containsKey(nums[i])) {
                indices[0] = hashMap.get(nums[i]);
                indices[1] = i;
                return indices;
            }
            hashMap.put(target - nums[i], i);
        }
        return indices;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] indices = solution.twoSum(new int[]{2, 7, 11, 15}, 9);
        
        for (int index : indices) {
            System.out.println(index);
        }
    }
}
