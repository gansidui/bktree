package bktree

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

func (this *BKTree) find(rt *bktreeNode, s string, k int, n int) (ret []string) {
	if n == 0 {
		return []string{}
	}

	d := Levenshtein(rt.str, s)
	if d <= k {
		ret = append(ret, rt.str)
		if n >= 0 && len(ret) >= n {
			return ret[0:n]
		}
	}

	dx, dy := max(0, d-k), d+k
	for i := dx; i <= dy; i++ {
		if rt.child[i] != nil {
			ret = append(ret, this.find(rt.child[i], s, k, n)...)
			if n >= 0 && len(ret) >= n {
				return ret[0:n]
			}
		}
	}
	return ret
}

// if n < 0, there is no limit on the number of find strings.
func (this *BKTree) Find(s string, k int, n int) []string {
	return this.find(this.root, s, k, n)
}

func (this *BKTree) Levenshtein(s1, s2 string) int {
	return Levenshtein(s1, s2)
}

func Levenshtein(s1, s2 string) int {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	m := len(runes1)
	n := len(runes2)

	// roll array
	d := make([][]int, 2)
	d[0] = make([]int, n+1)
	d[1] = make([]int, n+1)

	turn, pre := 0, 0
	for i := 0; i <= n; i++ {
		d[turn][i] = i
	}
	for i := 1; i <= m; i++ {
		pre = turn
		turn = (turn + 1) % 2
		d[turn][0] = i

		for j := 1; j <= n; j++ {
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
