package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// int型のx,yを使って、倍精度のfloatに変換しながら計算
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// 倍精度浮動小数点数の変数fを符号なし整数に変換してzに代入
	// var z uint = uint(f)
	// 以下のようにシンプルに記述する事も出来る。
	z := uint(f)
	fmt.Println(x, y, z)
}
