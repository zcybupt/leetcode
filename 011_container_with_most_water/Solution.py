class Solution:
    def maxArea(self, height: list) -> int:
        if not height or len(height) == 0: return 0

        l, r = 0, len(height) - 1
        ans, area = 0, 0

        while l < r:
            area = min(height[l], height[r]) * (r - l)
            ans = max(area, ans)
            if height[l] < height[r]: l += 1
            else: r -= 1

        return ans


if __name__ == '__main__':
    solution = Solution()
    print(solution.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))
