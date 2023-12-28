package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// 構造体のアドレスを代入し初期化
	p := &v
	// (*p).Xのように書く事で、フィールドにアクセス出来る。（しかし、この表記法は面倒）
	// Goではこれを、代わりにp.Xと書く事も出来る。
	p.X = 1e9
	// ポインタを使って書き換わった構造体のXと初期値のままのYが出力される。
	fmt.Println(v)
}
