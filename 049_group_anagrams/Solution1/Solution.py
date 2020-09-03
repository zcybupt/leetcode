class Solution:
    def groupAnagrams(self, strs: list) -> list:
        results = {}

        for s in strs:
            key = ''.join(sorted(s))
            if key in results:
                results[key].append(s)
            else:
                results[key] = [s]

        return list(results.values())


if __name__ == '__main__':
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    solution = Solution()
    print(solution.groupAnagrams(strs))

