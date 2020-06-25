package app

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	if len(beginWord) != len(endWord) {
		return [][]string{}
	}

	if beginWord == endWord {
		return [][]string{
			{beginWord},
		}
	}

	availableWord := make(map[string]bool)
	for _, w := range wordList {
		availableWord[w] = true
	}

	if !availableWord[endWord] {
		return [][]string{}
	}

	transformed := make(map[string][][]string)
	transformed[beginWord] = [][]string{{beginWord}}
	visited := make(map[string]bool)
	for len(transformed[endWord]) == 0 && len(availableWord) > 0 {
		newTransformed := make(map[string][][]string)
		toBeDeleteAvailableWord := make(map[string]bool)
		for word, paths := range transformed {
			visited[word] = true
			for i := 0; i < len(word); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					newWord := word[:i] + string(ch) + word[i+1:]
					if visited[newWord] || !availableWord[newWord] {
						continue
					}
					toBeDeleteAvailableWord[newWord] = true
					newTransformPath := newTransformed[newWord]
					for j := 0; j < len(paths); j++ {
						copyPath := append([]string{}, paths[j]...)
						copyPath = append(copyPath, newWord)
						newTransformPath = append(newTransformPath, copyPath)
					}
					newTransformed[newWord] = newTransformPath
				}
			}
		}

		for k := range toBeDeleteAvailableWord {
			delete(availableWord, k)
		}

		if len(newTransformed) == 0 {
			return [][]string{}
		}
		transformed = newTransformed
	}

	return transformed[endWord]
}
