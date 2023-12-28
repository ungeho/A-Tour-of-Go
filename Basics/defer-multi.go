package main

import "fmt"

func main() {
	fmt.Println("counting")

	// deferにより、10回分の処理がmain関数の終了時（return）まで遅延される。
	// deferへ渡した関数が複数ある場合、呼び出しはスタック(stack)される。
	// 呼び出し元の関数（main）がreturnする時、deferへ渡した関数はLIFO(last in first out)の順番で実行される。
	// その為、最後に実行されたものが最初に取り出され、forの中の処理とは逆の順で出力される。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
