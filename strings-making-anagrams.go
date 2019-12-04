package main

func makeAnagram(a string, b string) int32 {
	var deletions int32

	stringCountA := stringCount(a)
	stringCountB := stringCount(b)

	for i := 0; i < 26; i++ {
		deletions += absoluteValue(stringCountA[i] - stringCountB[i])
	}

	return deletions
}

func stringCount(x string) []int32 {
	stringCount := make([]int32, 26)
	for _, value := range x {
		stringCount[int(value)-97]++
	}

	return stringCount
}

func absoluteValue(x int32) int32 {
	if x < 0 {
		return -x
	}

	return x
}

func main() {}
