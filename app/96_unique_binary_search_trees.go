package app

/*NumTrees Catalan Number:
g(n)=g(0)*g(n-1)+g(1)g(n-2)+g(2)*g(n-3)+......+g(n-1)*g(0)
*/
func NumTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i <= n; i++ {
		for k := 1; k <= i; k++ {
			g[i] += g[k-1] * g[i-k]
		}
	}
	return g[n]
}
