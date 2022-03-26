## bktree

编辑距离（Edit Distance），又称Levenshtein距离，是指两个字串之间，由一个转成另一个所需的最少编辑操作次数。

许可的编辑操作包括将一个字符替换成另一个字符，插入一个字符，删除一个字符。


bktree的作用是：

给定一个词典（很多字符串），然后每输入一个字符串和一个数字k,

从词典中查找出与该字符串编辑距离小于等于k的字符串。

~~~ go
package main

import (
	"fmt"
	"github.com/gansidui/bktree"
	"log"
)

func main() {
	if bktree.Levenshtein("hello", "Aelo") != 2 {
		log.Fatal()
	}

	if bktree.Levenshtein("我爱你", "你爱我") != 2 {
		log.Fatal()
	}

	bk := bktree.New()
	bk.SetLevenshteinLimit(50)

	bk.Insert("ABCD")
	bk.Insert("ACED")
	bk.Insert("SBDE")

	ret := bk.Find("AABB", 3, 2)
	fmt.Println(ret)
}

~~~



## LICENSE

MIT