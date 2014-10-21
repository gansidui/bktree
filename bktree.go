package bktree

import (
	"unicode/utf8"
)

const DEFAULT_MAX_LEVENSHTEIN = 50

type bktreeNode struct {
	str   string
	child []*bktreeNode
}

func newBktreeNode(s string, limit int) *bktreeNode {
	return &bktreeNode{
		str:   s,
		child: make([]*bktreeNode, limit+1),
	}
}

type BKTree struct {
	root             *bktreeNode
	size             int
	levenshteinLimit int
}

func New() *BKTree {
	return &BKTree{
		root:             nil,
		size:             0,
		levenshteinLimit: DEFAULT_MAX_LEVENSHTEIN,
	}
}

func (this *BKTree) SetLevenshteinLimit(limit int) {
	this.levenshteinLimit = limit
}

func (this *BKTree) GetLevenshteinLimit() int {
	return this.levenshteinLimit
}

func (this *BKTree) Size() int {
	return this.size
}

func (this *BKTree) insert(rt *bktreeNode, s string) bool {
	d := Levenshtein(rt.str, s)
	if d > this.levenshteinLimit {
		return false
	}

	if rt.child[d] == nil {
		rt.child[d] = newBktreeNode(s, this.levenshteinLimit)
		return true
	} else {
		return this.insert(rt.child[d], s)
	}
}

func (this *BKTree) Insert(s string) bool {
	if this.root == nil {
		this.root = newBktreeNode(s, this.levenshteinLimit)
		this.size++
		return true
	}

	if this.insert(this.root, s) {
		this.size++
		return true
	}

	return false
}

func (this *BKTree) find(rt *bktreeNode, s string, k int) (ret []string) {
	d := Levenshtein(rt.str, s)
	if d <= k {
		ret = append(ret, rt.str)
	}

	dx, dy := max(0, d-k), d+k
	for i := dx; i <= dy; i++ {
		if rt.child[i] != nil {
			ret = append(ret, this.find(rt.child[i], s, k)...)
		}
	}
	return ret
}

func (this *BKTree) Find(s string, k int) []string {
	return this.find(this.root, s, k)
}

func Levenshtein(s1, s2 string) int {
	m, n := utf8.RuneCountInString(s1), utf8.RuneCountInString(s2)
	runes1, runes2 := make([]rune, m), make([]rune, n)

	// copy runes
	i, j := 0, 0
	for _, v := range s1 {
		runes1[i] = v
		i++
	}
	for _, v := range s2 {
		runes2[j] = v
		j++
	}

	// roll array
	d := make([][]int, 2)
	d[0] = make([]int, n+1)
	d[1] = make([]int, n+1)

	turn, pre := 0, 0
	for i = 0; i <= n; i++ {
		d[turn][i] = i
	}
	for i = 1; i <= m; i++ {
		pre = turn
		turn = (turn + 1) % 2
		d[turn][0] = i

		for j = 1; j <= n; j++ {
			if runes1[i-1] == runes2[j-1] {
				d[turn][j] = d[pre][j-1]
			} else {
				d[turn][j] = min(min(d[pre][j]+1, d[turn][j-1]+1), d[pre][j-1]+1)
			}
		}
	}

	return d[turn][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
