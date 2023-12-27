package main

import "fmt"

// 数値の定数は高精度な値
const (
	// 巨大な数、2^100
	Big = 1 << 100
	// (2^100 / 2^99)-> 2
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
