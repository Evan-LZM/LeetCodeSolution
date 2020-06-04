package app

func AddBinary(a string, b string) string {
	result := []byte{}
	sum := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || sum != 0; i, j = i-1, j-1 {
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		result = append(result, byte(sum)%2+'0')
		sum /= 2
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
