
def longestPalindrome(self, s):
    s = '#' + '#'.join(s) + '#'
    lens = len(s)
    p = [0] * lens
    mx = 0
    id = 0
    for i in range(lens):
        if mx > i:
            p[i] = min(mx-i, p[int(2*id-i)])
        else:
            p[i] = 1
        while i-p[i] >= 0 and i+p[i] < lens and s[i-p[i]] == s[i+p[i]]:
            p[i] += 1
        if(i+p[i]) > mx:
            mx, id = i+p[i], i
    i_res = p.index(max(p))
    s_res = s[i_res-(p[i_res]-1):i_res+p[i_res]]
    return s_res.replace('#', '')
