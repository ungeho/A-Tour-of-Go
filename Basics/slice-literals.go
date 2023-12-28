package main

import "fmt"

func main() {
	// [6]int{2, 3, 5, 7, 11, 13}と同様の配列リテラルを作成し
	// その配列リテラルを参照するスライスを作成する。
	// あくまで参照先の配列があり、スライス自体は配列を参照するだけで値を持っていない。
	// 長さのない配列リテラルのように扱える。
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// boolのスライス
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 構造体のスライス
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
