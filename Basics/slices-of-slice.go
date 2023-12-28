package main

import (
	"fmt"
	"strings"
)

func main() {

	// 三目並べのボードを作成
	// 文字列の二重スライスを定義し、スライスの中にスライスを入れている。
	// スライスは、他のスライスを含む任意の型を入れる事が出来る。
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 盤面の状況
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// 盤面を表示
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
