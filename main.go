package main

import (
	"bufio"
	"fmt"
	"github.com/AndrewVos/o"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("/Users/andrewvos/Desktop/all_categories")
	scanner := bufio.NewScanner(file)

	wordCount := 0
	l := NewLevenshteinDistance()
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "category: ", "", -1)
		words := strings.Split(line, " ")
		for _, word := range words {
			l.AddWord(word)
			wordCount += 1
		}
	}
	fmt.Println(wordCount, " words")

	for {
		var query string
		fmt.Println("Enter search query:")
		_, err := fmt.Scanln(&query)
		if err == nil {
			o.O(l.FindCloseWords(query, 1))
		}
	}
}
