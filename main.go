package main

import (
	"fmt"

	"github.com/landanqrew/goAdvent/osUtils"
	solutions2022 "github.com/landanqrew/goAdvent/solutions/2022"
	solutions2023 "github.com/landanqrew/goAdvent/solutions/2023"
)

func main() {
	dir := osUtils.ListDirectoryContents(".", true)
	if false {
		dir.Print()
	} 

	// day 1
	solutions2023.Day1()
	solutions2023.Day2(false, false)
	solutions2023.Day3(false, true)
	solutions2022.Day9(false)
	fmt.Println(osUtils.GetRegexMatches(`\S+`, "hello world!")[1])
}
/*
func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		_, ok := nameCounts[[]rune(name)[0]]
		if ok {
			_, ok := nameCounts[[]rune(name)[0]][name]
			if ok {
				nameCounts[[]rune(name)[0]][name]++
			} else {
				nameCounts[[]rune(name)[0]][name] = 1
			}
		} else {
			nameCounts[[]rune(name)[0]] = make(map[string]int)
			nameCounts[[]rune(name)[0]][name] = 1
		}
	}
	return nameCounts
}

func countDistinctWords(messages []string) int {
	wc := 0
	wordMap := make(map[string]bool)
	for _, m := range messages {
		lower := strings.ToLower(m)
		words := strings.Split(lower, " ")
		for _, word := range words {
			_, ok := wordMap[word]
			if !ok {
				wc += 1
				wordMap[word] = true
			}
		}
	}
	return wc
}
*/