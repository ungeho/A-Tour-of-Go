package main

import "fmt"

func main() {
	sum := 0
	//for ステートメントの初期化,条件式,後処理の三つの部分を括る括弧()は必要ない。
	// ただし、{}は常に必要
	for i := 0; i < 10; i++ {
		sum += i
	}
	// 下記の1行は構文エラーになる。{}を省略する事は出来ない。
	// for i := 0; i < 10; i++ sum+=i
	fmt.Println(sum)
}
