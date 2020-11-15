class Solution:
    def maxProfit(self, prices: list) -> int:
        cost, profit = float('inf'), 0
        for price in prices:
            cost = min(cost, price)
            profit = max(profit, price - cost)

        return profit


if __name__ == '__main__':
    prices = [7, 1, 5, 3, 6, 4]
    print(Solution().maxProfit(prices))
