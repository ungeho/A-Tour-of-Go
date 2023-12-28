package main

import "fmt"

func main() {
	// nilスライスを定義
	var s []int
	printSlice(s)

	// append(追加)はnilスライスに対しても機能する。
	// appendの戻り値は追加元のスライスと追加する変数群を合わせたスライス
	// もし、元の配列 s が変数軍を追加する際に容量が小さい場合は、より大きいサイズの配列を割り当て直し
	// 戻り値となるスライスは新しい割当先を示すようになる。
	// nilスライスの末尾に0の内容を追加する。
	s = append(s, 0)
	printSlice(s)

	// スライスの末尾に1の内容を追加する。
	s = append(s, 1)
	printSlice(s)

	// スライスの末尾に2,3,4の内容を追加する。
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
