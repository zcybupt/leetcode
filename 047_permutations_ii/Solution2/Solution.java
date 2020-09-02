import java.util.*;

class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> results = new ArrayList<>();
        Arrays.sort(nums);
        results.add(arrToList(nums));
        while (nextPermutation(nums)) {
            results.add(arrToList(nums));
        }

        return results;
    }

    public List<Integer> arrToList(int[] nums) {
        List<Integer> result = new ArrayList<>();
        for (int num : nums) result.add(num);
        return result;
    }

    public boolean nextPermutation(int[] nums) {
        int arrLen = nums.length;
        int i = arrLen - 2, j = arrLen - 1, k = arrLen - 1;

        while (i >= 0 && nums[i] >= nums[j]) {
            i--; j--;
        }

        if (i == -1) return false;

        while (nums[i] >= nums[k]) k--;
        swap(nums, i, k);

        for (i = j, j = arrLen - 1; i < j; i++, j--) {
            swap(nums, i, j);
        }

        return true;
    }

    public void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{3, 3, 0, 3};
        Solution solution = new Solution();
        List<List<Integer>> results = solution.permuteUnique(nums);
        for (List<Integer> result : results) {
            for (int ele : result) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}
