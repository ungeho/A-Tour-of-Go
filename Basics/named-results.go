package main

import "fmt"

// 戻り値に型だけでなく名前をつけている (x, y int)
// 戻り値に名前をつけると、関数の最初で定義した変数名として扱われる。
// 戻り値の意味を示す名前として、関数のドキュメントとして表現できる。
func split(sum int) (x, y int) {
	// 名前をつけた戻り値はreturnに何も書かずに返す事が出来る。
	// この書き方は、短い関数でのみの利用が推奨され、長い関数では読みやすさに悪影響が出る。
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
