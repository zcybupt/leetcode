class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        char_set = set()
        rk, ans = -1, 0
        length = len(s)

        for i in range(length):
            if i != 0: char_set.remove(s[i - 1])

            while rk + 1 < length and s[rk + 1] not in char_set:
                char_set.add(s[rk + 1])
                rk += 1
            ans = max(ans, rk - i + 1)

        return ans


if __name__ == '__main__':
    solution = Solution()
    print(solution.lengthOfLongestSubstring('abcabcbb'))
