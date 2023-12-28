package main

import (
	"fmt"
	"math"
)

// ニュートン法
func Sqrt(x float64, max int) float64 {
	z := float64(1.0)
	for i := 1; i <= max; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Printf("ループ回数\t出力結果\t相対誤差\n")
	for i := 1; i <= 10; i++ {
		fmt.Printf(
			"%02d:\t\t%e\t%e\n",
			i,
			Sqrt(2, i),
			math.Abs((Sqrt(2, i)-math.Sqrt(2))/math.Sqrt(2)),
		)
	}
}
