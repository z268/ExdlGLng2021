package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordStat struct {
	word    string
	order   int
	repeats int
}

type Stats []WordStat

func (stats Stats) asStrings() []string {
	res := make([]string, 0, len(stats))
	for _, stat := range stats {
		res = append(res, fmt.Sprintf("%v(%v)", stat.word, stat.repeats))
	}
	return res
}

func (stats Stats) sort() {
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].order < stats[j].order
	})
}

func getWords(input string) []string {
	input = strings.ToLower(input)
	r := regexp.MustCompile(`[\wа-я]+`)
	return r.FindAllString(input, -1)
}

func getStats(words []string) Stats {
	statsByWord := map[string]WordStat{}
	for idx, word := range words {
		if wordStat, ok := statsByWord[word]; !ok {
			statsByWord[word] = WordStat{word, idx, 1}
		} else {
			wordStat.repeats += 1
			wordStat.order -= len(words)
			statsByWord[word] = wordStat
		}
	}

	stats := make(Stats, 0, len(statsByWord))
	for _, wordStat := range statsByWord {
		stats = append(stats, wordStat)
	}

	return stats
}

func analyzeString(input string) string {
	words := getWords(input)
	stats := getStats(words)
	stats.sort()
	return strings.Join(stats.asStrings(), " ")
}

func main() {
	fmt.Print("Enter string: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(analyzeString(input))
}
