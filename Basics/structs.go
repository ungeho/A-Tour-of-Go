package main

import "fmt"

// 構造体
type Vertex struct {
	// フィールド
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
