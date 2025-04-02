package stringprograms

import (
	"sort"
)

func IsAnagrams1(source, target string) bool {
	if len(source) != len(target) {
		return false
	}

	sourceArray := []rune(source)
	sort.Slice([]rune(sourceArray), func(i, j int) bool {
		return sourceArray[i] < sourceArray[j]
	})

	targetArray := []rune(target)
	sort.Slice([]rune(targetArray), func(i, j int) bool {
		return targetArray[i] < targetArray[j]
	})
	for i := 0; i < len(sourceArray); i++ {
		if sourceArray[i] != targetArray[i] {
			return false
		}
	}
	return true
}

func IsAnagrams2(source, target string) bool {
	if len(source) != len(target) {
		return false
	}
	var sourceCount [26]int
	var targetCount [26]int
	for i := 0; i < len(source); i++ {
		sourceCount[source[i]-'a']++
		targetCount[target[i]-'a']++
	}
	for i := 0; i < len(sourceCount); i++ {
		if sourceCount[i] != targetCount[i] {
			return false
		}
	}
	return true
}
func IsAnagrams3(source, target string) bool {

	if len(source) != len(target) {
		return false
	}

	var alphabetCounter [26]int
	for i := 0; i < len(source); i++ {
		alphabetCounter[source[i]-'a']++
		alphabetCounter[target[i]-'a']--
	}

	for j := 0; j < len(alphabetCounter); j++ {
		if alphabetCounter[j] != 0 {
			return false
		}
	}
	return true
}

func IsAnagrams4(source, target string) bool {

	if len(source) != len(target) {
		return false
	}
	sourceCount := make(map[rune]int)
	targetCount := make(map[rune]int)

	for _, rune := range source {
		sourceCount[rune]++
	}

	for _, rune := range target {
		targetCount[rune]++
	}

	for letter, sourceCountValue  := range sourceCount {
		if targetCountValue, ok := targetCount[letter]; !ok || sourcecount != targetCount {
			return false
		}
	}
	return true
}
