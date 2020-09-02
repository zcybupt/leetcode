class Solution:
    def permuteUnique(self, nums: list) -> list:
        results = []
        arr_len = len(nums)

        def next_permutation() -> bool:
            i, j, k = arr_len - 2, arr_len - 1, arr_len - 1

            while i >= 0 and nums[i] >= nums[j]:
                i -= 1
                j -= 1

            if i == -1: return False

            while nums[i] >= nums[k]: k -= 1
            nums[i], nums[k] = nums[k], nums[i]

            i, j = j, arr_len - 1
            while i < j:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j -= 1

            return True

        nums.sort()
        results.append(nums[:])
        while next_permutation():
            results.append(nums[:])

        return results


if __name__ == '__main__':
    nums = [3, 0, 3, 3]
    solution = Solution()
    results = solution.permuteUnique(nums)
    for result in results:
        print(result)

