package main

// Go言語の基本型(組み込み方)は以下の通り
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte (uint8の別名)
// rune (int32の別名)
// (Unicode のコードポイントを表す)
// float32 float64
// complex64 complex128

import (
	"fmt"
	"math/cmplx"
)

var (
	// bool型 falseを代入
	ToBe bool = false
	// 符号なし64bit 最大値(2^64 - 1)を代入
	MaxInt uint64 = 1<<64 - 1
	// 複素数
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// 型と値を表示
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
