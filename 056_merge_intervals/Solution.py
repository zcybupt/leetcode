class Solution:
    def quick_sort(self, nums: list, l: int, r: int):
        if l < r:
            p_index = self.partition(nums, l, r)
            self.quick_sort(nums, l, p_index - 1)
            self.quick_sort(nums, p_index + 1, r)

    def partition(self, nums: list, l: int, r: int) -> int:
        p = nums[l]
        while l != r:
            while l < r and nums[r][0] >= p[0]: r -= 1
            nums[l] = nums[r]
            while l < r and nums[l][0] <= p[0]: l += 1
            nums[r] = nums[l]
        nums[l] = p

        return l

    def merge(self, intervals: list) -> list:
        if not intervals or len(intervals) <= 1: return intervals

        # self.quick_sort(intervals, 0, len(intervals) - 1)
        intervals.sort(key=lambda x: x[0])
        result = [intervals[0]]
        for interval in intervals[1:]:
            if result[-1][1] < interval[0]: result.append(interval)
            else: result[-1][1] = max(result[-1][1], interval[1])

        return result


if __name__ == '__main__':
    nums = [[2, 6], [1, 3], [9, 13], [15, 18], [8, 10]]
    print(Solution().merge(nums))

