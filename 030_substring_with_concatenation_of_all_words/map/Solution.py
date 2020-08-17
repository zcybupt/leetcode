class Solution:
    def findSubstring(self, s: str, words: list) -> list:
        results = []
        s_len = len(s)
        arr_len = len(words)
        if s_len == 0 or arr_len == 0: return results
        word_len = len(words[0])

        word_map = {}
        for word in words:
            value = word_map.get(word)
            word_map[word] = value + 1 if value else 1

        for i in range(s_len - word_len * arr_len + 1):
            cur_map = {}
            count = 0
            while count < arr_len:
                cur_word = s[i + count * word_len: i + (count + 1) * word_len]
                if cur_word in word_map.keys():
                    value = cur_map.get(cur_word)
                    if value and value + 1 > word_map.get(cur_word): break
                    cur_map[cur_word] = value + 1 if value else 1
                else: break
                count += 1
            if count == arr_len: results.append(i)

        return results


if __name__ == '__main__':
    s = 'barfoothefoobarman'
    words = ['foo', 'bar']
    solution = Solution()
    print(solution.findSubstring(s, words))

