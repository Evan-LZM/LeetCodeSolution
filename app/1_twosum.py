class Solution(object):
    def twoSum(self, nums, target):
        d = {}  # assert the dictionary type to store value and key
        for i, num in enumerate(nums):
            if target-num in d:
                return [d[target-num], i]
            d[num] = i


nums = [2, 7, 11, 15]
target = 18
s = Solution()
print(s.twoSum(nums, target))
