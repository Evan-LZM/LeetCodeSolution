package app

type trieNode struct {
	isWord   bool
	children [26]*trieNode
	word     string
}

//FindWords using trie structure to store
func FindWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(words[0]) == 0 {
		return []string{}
	}
	//create root node to store all of the words in tree structure
	root := &trieNode{children: [26]*trieNode{}, word: ""}
	for _, word := range words {
		cur := root
		for i := 0; i < len(word); i++ {
			if cur.children[int(word[i]-'a')] == nil {
				cur.children[int(word[i]-'a')] = &trieNode{
					children: [26]*trieNode{},
					word:     cur.word + word[i:i+1],
				}
			}
			cur = cur.children[int(word[i]-'a')]
		}
		//right now the cur is leaves
		cur.isWord = true
	}
	m := make(map[string]bool)
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			//check every node in the board
			helperSearchWord(board, root, visited, i, j, m)
		}
	}
	res := []string{}
	for k := range m {
		res = append(res, k)
	}
	return res
}

// var directions = [][]int{
// 	[]int{1, 0},
// 	[]int{-1, 0},
// 	[]int{0, 1},
// 	[]int{0, -1},
// }

func helperSearchWord(board [][]byte, node *trieNode, visited [][]bool, i, j int, m map[string]bool) {
	if node.isWord {
		m[node.word] = true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] || node.children[int(board[i][j]-'a')] == nil {
		return
	}
	visited[i][j] = true
	for _, d := range directions {
		ii, jj := i+d[0], j+d[1]
		helperSearchWord(board, node.children[int(board[i][j]-'a')], visited, ii, jj, m)
	}
	visited[i][j] = false

}
