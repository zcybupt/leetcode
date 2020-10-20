class Solution:
    def majorityElement(self, nums: list) -> int:
        votes, x = 0, 0
        for num in nums:
            if votes == 0: x = num
            if x == num: votes += 1
            else: votes -= 1

        return x


if __name__ == '__main__':
    nums = [1, 2, 2, 3, 2, 3, 2]
    print(Solution().majorityElement(nums))

