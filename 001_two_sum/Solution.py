class Solution:
    def twoSum(self, nums: list, target: int) -> list:
        if not nums or len(nums) == 0: return

        for i in range(len(nums)):
            res = target - nums[i]
            index = nums.index(res) if res in nums else -1
            if index != -1 and index != i:
                return [i, index]

    def twoSum2(self, nums: list, target: int) -> list:
        if not nums or len(nums) == 0: return

        res_dict = dict()
        for i in range(len(nums)):
            index = res_dict.get(nums[i])
            if index is not None:
                return [index, i]
            res_dict[target - nums[i]] = i


if __name__ == '__main__':
    solution = Solution()
    print(solution.twoSum2([3, 7, 11, 15], 10))
