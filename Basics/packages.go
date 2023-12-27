package main

// 入出力と数学の乱数
import (
	"fmt"
	"math/rand"
)

func main() {
	// 文字列と0~9のランダムな数字
	fmt.Println("My facorite number is", rand.Intn(10))
}
