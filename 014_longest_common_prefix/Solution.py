class Solution:
    def longestCommonPrefix(self, strs: list) -> str:
        if not strs or len(strs) == 0: return ''

        common_prefix = ''
        for tmp in zip(*strs):
            tmp = set(tmp)
            if len(tmp) == 1: common_prefix += tmp.pop()
            else: break

        return common_prefix


if __name__ == '__main__':
    solution = Solution()
    test = ['flower', 'flour', 'flow']
    print(solution.longestCommonPrefix(test))
