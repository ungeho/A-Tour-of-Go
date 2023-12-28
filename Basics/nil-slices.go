package main

import "fmt"

func main() {
	var s []int
	// nilスライスは長さも容量も0であり、元となる配列も持っていない。
	fmt.Println(s, len(s), cap(s))
	// スライスのゼロ値はnil（ポインタと同じ）
	if s == nil {
		fmt.Println("nil!")
	}
}
