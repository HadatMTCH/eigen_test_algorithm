package main

import (
	"fmt"
	"strings"
)

// 1. Reverse string with number at the end
func reverseWithNumber(s string) string {
	// Find the position of the first digit
	digitPos := strings.IndexFunc(s, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	if digitPos == -1 {
		// If no digit found, reverse the whole string
		return reverseString(s)
	}

	// Reverse only the part before the digit
	reversed := reverseString(s[:digitPos])
	return reversed + s[digitPos:]
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 2. Find the longest word in a sentence
func longestWord(sentence string) string {
	words := strings.Fields(sentence)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// 3. Count occurrences of query words in input array
func countOccurrences(input, query []string) []int {
	countMap := make(map[string]int)
	for _, word := range input {
		countMap[word]++
	}

	result := make([]int, len(query))
	for i, q := range query {
		result[i] = countMap[q]
	}
	return result
}

// 4. Calculate diagonal difference in a matrix
func diagonalDifference(matrix [][]int) int {
	n := len(matrix)
	primarySum, secondarySum := 0, 0

	for i := 0; i < n; i++ {
		primarySum += matrix[i][i]
		secondarySum += matrix[i][n-1-i]
	}

	diff := primarySum - secondarySum
	if diff < 0 {
		return -diff
	}
	return diff
}

func main() {
	// Test the functions
	fmt.Println(reverseWithNumber("NEGIE1"))
	fmt.Println(longestWord("Saya sangat senang mengerjakan soal algoritma"))
	fmt.Println(countOccurrences([]string{"xc", "dz", "bbb", "dz"}, []string{"bbb", "ac", "dz"}))
	fmt.Println(diagonalDifference([][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}))
}
