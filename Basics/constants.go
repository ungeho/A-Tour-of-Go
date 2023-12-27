package main

import "fmt"

const Pi = 3.14

func main() {
	// 定数(constant)は、constキーワードを使って変数と同じように宣言
	// 定数は、文字(character)、文字列(string)、boolean（論理型）、数値（numeric）のみで使える。
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
