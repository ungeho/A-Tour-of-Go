package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	// Name: ...フィールドの一部だけを列挙する事が出来る（この方法では、フィールドの指定順序は関係ない）
	v2 = Vertex{X: 1}  // X: 1として、Xのみを1で初期化,Yの値は暗黙的に0になる。
	v3 = Vertex{}      // 暗黙的に X:0 と Y:0
	p  = &Vertex{1, 2} // 新しく割り当てられたstructへのポインタを戻す。
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
