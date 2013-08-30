package main

type LevenshteinDistance struct {
	Trie *TrieNode
}

func NewLevenshteinDistance() *LevenshteinDistance {
	t := NewTrieNode()
	return &LevenshteinDistance{Trie: t}
}

func (l *LevenshteinDistance) AddWord(word string) {
	l.Trie.Insert(word)
}

type LevenshteinWordMatch struct {
	Word     string
	Distance int
}

var results []LevenshteinWordMatch

func minInt(mins ...int) int {
	smallest := mins[0]
	for _, i := range mins {
		if i < smallest {
			smallest = i
		}
	}
	return smallest
}

func (l *LevenshteinDistance) recurse(node *TrieNode, letter byte, word string, previousRow []int, maximumDistance int) {
	columns := len(word) + 1
	currentRow := []int{previousRow[0] + 1}

	wordBytes := []byte(word)

	for column := 1; column < columns; column++ {
		var insertCost int
		if column == 0 {
			insertCost = currentRow[len(currentRow)-1]
		} else {
			insertCost = currentRow[column-1] + 1
		}

		deleteCost := previousRow[column] + 1
		var replaceCost int
		if wordBytes[column-1] != letter {
			replaceCost = previousRow[column-1] + 1
		} else {
			replaceCost = previousRow[column-1]
		}
		currentRow = append(currentRow, minInt(insertCost, deleteCost, replaceCost))
	}

	if currentRow[len(currentRow)-1] <= maximumDistance && node.Word != "" {
		results = append(results, LevenshteinWordMatch{Word: node.Word, Distance: currentRow[len(currentRow)-1]})
	}

	if minInt(currentRow...) <= maximumDistance {
		for childLetter, childNode := range node.Children {
			l.recurse(childNode, childLetter, word, currentRow, maximumDistance)
		}
	}
}

func (l *LevenshteinDistance) FindCloseWords(word string, maximumDistance int) []LevenshteinWordMatch {
	results = []LevenshteinWordMatch{}
	currentRow := []int{}
	for i := 0; i <= len(word); i++ {
		currentRow = append(currentRow, i)
	}

	for letter, node := range l.Trie.Children {
		l.recurse(node, letter, word, currentRow, maximumDistance)
	}
	return results
}
