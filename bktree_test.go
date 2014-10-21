package bktree

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	s1, s2 := "hello", "world"
	if Levenshtein(s1, s2) != 4 {
		t.Fatal()
	}

	if Levenshtein("", "AA") != 2 {
		t.Fatal()
	}

	if Levenshtein("a", "b") != 1 || Levenshtein("aa", "bb") != 2 {
		t.Fatal()
	}

	if Levenshtein("abcd", "bcde") != 2 {
		t.Fatal()
	}

	if Levenshtein("AB好C", "你好啊") != 3 {
		t.Fatal()
	}

	if Levenshtein("你好世界啊", "世界啊你好") != 4 {
		t.Fatal()
	}
}
