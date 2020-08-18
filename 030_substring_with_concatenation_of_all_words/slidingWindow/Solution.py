class Solution:
    def findSubstring(self, s: str, words: list) -> list:
        results = []
        s_len, arr_len = len(s), len(words)
        if s_len == 0 or arr_len == 0: return results
        word_len = len(words[0])

        word_map = {}
        for word in words:
            word_map[word] = word_map.get(word, 0) + 1

        cur_map = {}
        window_size = word_len * arr_len
        for i in range(word_len):
            start = i
            while start + window_size <= s_len:
                cur_str = s[start: start + window_size]
                cur_map.clear()
                j = arr_len
                while j > 0:
                    word = cur_str[(j - 1) * word_len: j * word_len]
                    count = cur_map.get(word, 0) + 1
                    if count > word_map.get(word, 0): break
                    cur_map[word] = count
                    j -= 1
                if j == 0: results.append(start)
                start += word_len * max(j, 1)

        return results


if __name__ == '__main__':
    s = 'barfoothefoobarman'
    words = ["foo","bar"]
    solution = Solution()
    print(solution.findSubstring(s, words))

