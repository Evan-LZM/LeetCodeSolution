package app

func IsScramble(s1, s2 string) bool {
	A := make([]int, 256)
	if len(s1) != len(s2) {
		return false
	}
	length := len(s1)
	if length == 0 || s1 == s2 {
		return true
	}
	//compare s1 and s2 whther have the same character
	for _, c := range []byte(s1) {
		A[c]++
	}
	for _, c := range []byte(s2) {
		A[c]--
	}
	for _, c := range A {
		if c != 0 {
			return false
		}
	}
	for i := 1; i < length; i++ {
		if IsScramble(s1[0:i], s2[0:i]) && IsScramble(s1[i:], s2[i:]) {
			return true
		}
		if IsScramble(s1[0:i], s2[length-i:]) && IsScramble(s1[i:], s2[:length-i]) {
			return true
		}
	}
	return false
}
