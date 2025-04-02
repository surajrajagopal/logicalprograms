package stringprograms

func RemoveDuplicates(str string) map[string]int {
	frequencyCount := make(map[string]int)
	for _, v := range str {
		if _, ok := frequencyCount[string(v)]; !ok {
			frequencyCount[string(v)] = 1
		} else {
			frequencyCount[string(v)]++
		}
	}

	CopiedMap := make(map[string]int)
	for k, v := range frequencyCount {
		if v == 1 {
			CopiedMap[k] = v
		}
	}
	return CopiedMap
}

func removeduplicatefromslice(fruits []string) []string {
	uniqueMap := make(map[string]bool)
	str := []string{}
	for _, v := range fruits {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			str = append(str, v)
		}
	}
	return str
}

