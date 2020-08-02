class Solution:
    def preprocess(self, s: str) -> str:
        result = '^'
        for ele in s:
            result += '#' + ele
        return result + '#$'

    def longestPalindrome(self, s: str) -> str:
        tmp_str = self.preprocess(s)
        str_len = len(tmp_str)
        p = [0] * str_len
        c, r = 0, 0
        index = 0
        max_len = 0

        for i in range(1, str_len - 1):
            p[i] = min(p[2 * c - i], r - i) if r > i else 1

            while tmp_str[i + p[i]] == tmp_str[i - p[i]]:
                p[i] += 1

            if r < i + p[i]:
                r = i + p[i]
                c = i

            if max_len < p[i] - 1:
                max_len = p[i] - 1
                index = i

        start = (index - max_len) // 2
        return s[start: start + max_len]


if __name__ == '__main__':
    solution = Solution()
    print(solution.longestPalindrome('babad'))
