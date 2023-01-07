package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World, This is my first app, This language is awesome for app development."
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ReplaceAll(s, ",", " ")
	s = strings.ReplaceAll(s, "  ", " ")

	words := strings.Fields(s)
	wordsCountMap := map[string]int{}

	// for _, word := range words {
	// 	_, ok := wordsCountMap[word]
	// 	if ok {
	// 		wordsCountMap[strings.ToLower(word)]++
	// 	} else {
	// 		wordsCountMap[strings.ToLower(word)] = 1
	// 	}
	// }

	for _, word := range words {
		wordsCountMap[strings.ToLower(word)]++
	}

	fmt.Printf("%v", wordsCountMap)
}
