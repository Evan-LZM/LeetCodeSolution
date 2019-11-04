class Solution:
    def maxArea(self, height: List[int]) -> int:
        left, right, maxarea = 0, len(height)-1, 0
        while left < right:
            maxarea = max(maxarea, min(
                height[right], height[left])*(right-left))
            if height[left] < height[right]:
                left += 1
            else:
                right += 1
        return maxarea
