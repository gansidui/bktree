package bktree

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"testing"
	"time"
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

func TestInsert(t *testing.T) {
	bk := New()
	bk.SetLevenshteinLimit(2)

	if !bk.Insert("") {
		t.Fatal()
	}

	if bk.Insert("ABC") {
		t.Fatal()
	}

	if !bk.Insert("AB") {
		t.Fatal()
	}
}

func TestInsertAndFind(t *testing.T) {
	bk := New()
	bk.SetLevenshteinLimit(50)

	if bk.GetLevenshteinLimit() != 50 {
		t.Fatal()
	}

	bk.Insert("656")
	bk.Insert("67")
	bk.Insert("9313")
	bk.Insert("1178")
	bk.Insert("38")

	if bk.Size() != 5 {
		t.Fatal()
	}

	ret := bk.Find("87", 2, 2)
	if ret[0] != "67" || ret[1] != "38" {
		t.Fatal()
	}

	ret = bk.Find("87", 2, 1)
	if len(ret) != 1 {
		t.Fatal()
	}

	ret = bk.Find("87", 4, -1)
	if len(ret) != bk.Size() {
		t.Fatal()
	}
}

func Test(t *testing.T) {
	bk := New()
	bk.SetLevenshteinLimit(20)

	start := time.Now()
	for i := 0; i < 100000; i++ {
		buf := make([]byte, 5)
		io.ReadFull(rand.Reader, buf)
		s := base64.StdEncoding.EncodeToString(buf)
		bk.Insert(s)
	}
	fmt.Println("Insert", time.Since(start))

	start = time.Now()
	for i := 0; i < 10; i++ {
		buf := make([]byte, 5)
		io.ReadFull(rand.Reader, buf)
		s := base64.StdEncoding.EncodeToString(buf)
		ret := bk.Find(s, 5, 2)
		if len(ret) > 2 {
			t.Fatal()
		}
	}
	fmt.Println("Find", time.Since(start))
}
