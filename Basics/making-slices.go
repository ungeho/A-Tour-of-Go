package main

import "fmt"

func main() {
	// 組み込み関数(make)を使用してスライスを作成する事も出来る。
	// 動的サイズの配列を作成する方法
	// make関数はゼロ化された配列を割り当て、その配列を指すスライスを返す。

	// len(a)=5 つまり長さ5のゼロ化された配列のスライスを作成
	a := make([]int, 5)
	printSlice("a", a)

	// スライスの3番目の引数には、スライスの容量（capacity）を指定できる。
	// 長さ len(b)=0 、容量 cap(b)=5 のスライスを作成する。
	b := make([]int, 0, 5)
	printSlice("b", b)

	// bのスライスを元に、参照する要素番号を0~1にしたスライスを作成（len(c)=2）
	c := b[:2]
	printSlice("c", c)

	// cのスライスを元に、参照する要素番号を2~4にしたスライスを作成（len(d)=3 cap(d)=3）
	// また、b,c,dは全て同じ配列を指し示している。
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
