class Solution {
    public int trap(int[] heightArr) {
        if (heightArr == null || heightArr.length < 3) return 0;

        int arrLen = heightArr.length;
        int leftMax = heightArr[0];
        int rightMax = heightArr[arrLen - 1];
        int result = 0;
        int L = 1, R = arrLen - 2;

        while (L <= R) {
            if (leftMax <= rightMax) {
                result += Math.max(0, leftMax - heightArr[L]);
                leftMax = Math.max(leftMax, heightArr[L++]);
            } else {
                result += Math.max(0, rightMax - heightArr[R]);
                rightMax = Math.max(rightMax, heightArr[R--]);
            }
        }

        return result;
    }

    public static void main(String[] args) {
        int[] heightArr = new int[]{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
        Solution solution = new Solution();
        System.out.println(solution.trap(heightArr));
    }
}
