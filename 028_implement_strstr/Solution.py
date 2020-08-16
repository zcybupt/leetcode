class Solution:
    def get_next_table(self, needle):
        needle_len = len(needle)
        next_table = [0] * needle_len
        next_table[0] = -1
        lp, rp = -1, 0

        while rp < needle_len - 1:
            if needle[lp] == needle[rp] or lp == -1:
                lp += 1
                rp += 1
                next_table[rp] = lp
            else:
                lp = next_table[lp]

        return next_table

    def strStr(self, haystack, needle):
        haystack_len = len(haystack)
        needle_len = len(needle)
        if needle_len == 0: return 0

        haystack_cur, needle_cur = 0, 0
        next_table = self.get_next_table(needle)

        while haystack_cur < haystack_len and needle_cur < needle_len:
            if needle[needle_cur] == haystack[haystack_cur]:
                needle_cur += 1
                haystack_cur += 1
            else:
                if needle_cur == 0: haystack_cur += 1
                else: needle_cur = next_table[needle_cur]
        return haystack_cur - needle_cur if needle_cur == needle_len else -1


if __name__ == '__main__':
    solution = Solution()
    print(solution.strStr('hello', 'll'))
