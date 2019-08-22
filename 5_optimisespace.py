def longestPalindrome(self, s):
    k = len(s)
    olist = [0]*k
    nlist = [0]*k
    longestSubstr = ""
    longestLen = 0
    for j in range(0, k):
        for i in range(0, j+1):
            if j-i <= 1:
                if s[i] == s[j]:
                    nlist[i] = 1
                    len_t = j-i+1
                    if longestLen < len_t:
                        longestSubstr = s[i:j+1]
                        longestLen = len_t
            else:
                if s[i] == s[j] and olist[i+1]:
                    nlist[i] = 1
                    len_t = j-i+1
                    if longestLen < len_t:
                        longestSubstr = s[i:j+1]
                        longestLen = len_t
        olist = nlist
        nlist = [0]*k
    return longestSubstr
