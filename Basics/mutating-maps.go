package main

import "fmt"

// map m の操作
// mへ要素(elem)の挿入や更新:
// m[key] = elem
// 要素の取得:
// elem = m[key]
// 要素の削除:
// delete(m, key)
// キーに対する要素が存在するかどうか
// elem, ok = m[key]
// mにkeyがあれば、okはtrue
// mにkeyがなければ、okはfalse
// mapにkeyが存在しない場合、elemはmapの要素の型のゼロ値になる。

func main() {
	m := make(map[string]int)

	// mapのkey="Answer"に42を割り当てる。
	m["Answer"] = 42
	// key="Answer"に割り当てられた42が表示される。
	fmt.Println("The value:", m["Answer"])

	// mapのkey="Answer"に48を割り当てる。
	m["Answer"] = 48
	// key="Answer"に割り当てられた48が表示される。
	fmt.Println("The value:", m["Answer"])

	// mapのkey="Answer"の要素を削除する
	delete(m, "Answer")
	// 要素が削除されたため、0が表示される。
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が v に格納される。
	// キーに対する要素が存在するかどうがかがboolで ok に格納される。
	v, ok := m["Answer"]
	// 要素が削除された後な為、v=0でok=falseとなる。
	fmt.Println("The value:", v, "Present?", ok)
}
