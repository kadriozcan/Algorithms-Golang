package main

import (
	"fmt"
)

func main() {
	// Question 1 Example:
	wordList := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	sortingByNofA(wordList)

	// Question 2 Example:
	wordList2 := []string{"apple", "pie", "apple", "red", "red", "red"}
	mostRepeated(wordList2)
}

// Qusetion 1 (selection sort)
func sortingByNofA(list []string) {
	for i := 0; i < len(list); i++ {
		maxAId := i
		for j := i + 1; j < len(list); j++ {
			if nOfAs(list[maxAId]) < nOfAs(list[j]) {
				maxAId = j
			}
		}
		list[i], list[maxAId] = list[maxAId], list[i]
	}
	fmt.Println("Question 1 [Output]:", list)
}

func nOfAs(word string) int {
	count := 0
	for _, c := range word {
		if c == 'a' {
			count++
		}
	}
	return count
}

// ****************** End of Question 1 **************

// Question 2

// ****************** End of Question 2 **************

// Question 3
func mostRepeated(list []string) {
	words := make(map[string]int)
	var mostRepeatedWord string
	maxCount := 0

	for _, word := range list {
		words[word]++
		if words[word] > maxCount {
			maxCount = words[word]
			mostRepeatedWord = word
		}
	}
	fmt.Println("Queston 2 - Most repeated word:", mostRepeatedWord)
}

// ****************** End of Question 3 **************
