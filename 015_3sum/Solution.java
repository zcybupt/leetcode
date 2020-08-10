import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        int listLen = nums.length;
        List<List<Integer>> results = new ArrayList<>();
        if (listLen < 3) return results;

        Arrays.sort(nums);
        int L, R;

        for (int i = 0; i < listLen - 2; i++) {
            if (nums[i] > 0) return results;
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            L = i + 1;
            R = listLen - 1;
            while (L < R) {
                if (nums[i] + nums[L] + nums[R] == 0) {
                    results.add(new ArrayList<>(Arrays.asList(nums[i], nums[L], nums[R])));
                    while (L < R && nums[L] == nums[L + 1]) L++;
                    while (L < R && nums[R] == nums[R - 1]) R--;
                    L++; R--;
                } else if (nums[i] + nums[L] + nums[R] > 0) {
                    R--;
                } else {
                    L++;
                }
            }
        }

        return results;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{-1, 0, 1, 2, -1, -4};
        Solution solution = new Solution();
        List<List<Integer>> results = solution.threeSum(nums);

        for (List<Integer> result : results) {
            for (int ele : result) {
                System.out.print(ele + " ");
            }
            System.out.println();
        }
    }
}
