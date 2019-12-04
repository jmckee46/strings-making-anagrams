package main

import (
	"testing"
)

func TestStringsMakingAnagrams(t *testing.T) {
	a := "cde"
	b := "abc"

	deletions := makeAnagram(a, b)

	if deletions != 4 {
		t.Errorf("got %d instead of 4", deletions)
	}
}
