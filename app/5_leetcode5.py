def longestPalindrome(self, s):
    k = len(s)
    matrix = [[0 for i in range(k)] for i in range(k)]
    longestSubstr = ""
    longestLen = 0
    for j in range(0, k):
        for i in range(0, j+1):
            if j-i <= 1:
                if s[i] == s[j]:
                    matrix[i][j] = 1
                    if longestLen < j-i+1:
                        longestSubstr = s[i:j+1]
                        longestLen = j-i+1
            else:
                if s[i] == s[j] and matrix[i+1][j-1]:
                    matrix[i][j] = 1
                    if longestLen < j-i+1:
                        longestSubstr = s[i:j+1]
                        longestLen = j-i+1
    return longestSubstr
