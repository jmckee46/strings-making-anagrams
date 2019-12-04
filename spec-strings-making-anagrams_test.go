package main

import (
	"fmt"
	"testing"
)

func TestStringsMakingAnagrams(t *testing.T) {
	a := "cde"
	b := "abc"

	deletions := makeAnagram(a, b)

	fmt.Println(deletions)

	if deletions != 4 {
		t.Errorf("got %d instead of 4", deletions)
	}
}
