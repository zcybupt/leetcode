class Solution:
    def threeSumClosest(self, nums: list, target: int) -> int:
        list_len = len(nums)
        if not nums or list_len < 3: return 0

        nums.sort()
        min_diff = float('inf')
        result = 0

        for i in range(list_len):
            l = i + 1
            r = list_len - 1
            while l < r:
                cur_sum = nums[i] + nums[l] + nums[r]
                if cur_sum == target: return cur_sum
                cur_diff = abs(cur_sum - target)
                if cur_diff < min_diff:
                    result = cur_sum
                    min_diff = cur_diff

                if cur_sum > target: r -= 1
                else: l += 1

        return result


if __name__ == '__main__':
    solution = Solution()
    nums = [0, 1, 2]
    target = 0
    print(solution.threeSumClosest(nums, target))
