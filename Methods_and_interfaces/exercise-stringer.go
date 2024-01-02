package main

import "fmt"

type IPAddr [4]byte

//  IPAddr型の引数をドットで4つに区切った( dotted quad )表現で出力
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	// map
	// キーはstring型
	// 内容の部分はIPAddr型
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
