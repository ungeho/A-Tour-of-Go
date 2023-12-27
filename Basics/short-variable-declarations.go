package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数の中であれば、:=による暗黙的な型宣言が可能
	// 関数の外では、キーワードから始まる宣言(var,funcなど)が必要で、:=を使った暗黙的な宣言は利用出来ない。
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
