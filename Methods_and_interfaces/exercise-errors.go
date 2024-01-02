package main

import (
	"fmt"
	"math"
)

// エラー用の型
type ErrNegativeSqrt float64

// マシンイプシロン
// NextAfter(x,y)の時、yの方向に向かってxの後にくる次の表現可能な倍精度浮動小数点値を計算する。
// x=2,y=1の時、2から1に向かって次に表現可能な倍精度浮動小数点数値になるので
// 2-ε (ε=表現可能な倍精度浮動小数点数値の最低値)
// x=1,y=1の時,1から1に向かうので
// 1
// x=1,y=2の時、1から2に向かって次に表現可能な倍精度浮動小数点数値になるので
// 1+ε
// 1を引く事でεが得られる。
var eps float64 = math.Nextafter(1, 2) - 1

func Sqrt(x float64) (float64, error) {
	// 変数xが負数の時にエラーメッセージを表示する為
	// 変数xの値と、ErrNegativeSqrt型にキャストした変数xの値を返す。
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	// ニュートン法
	z := 1.0
	prev_z := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		// 前の項と今の項の絶対誤差がマシンイプシロン以下になったらループを抜ける。
		if math.Abs(z-prev_z) <= eps {
			break
		}
		prev_z = z
	}
	// 成功時はnilを返す。
	return z, nil
}

// エラーインターフェース
func (e ErrNegativeSqrt) Error() string {
	//エラーメッセージ、平方根のxには負数は無効である事を伝える。
	// ErrNegativeSqrt型を倍精度にキャストする。
	// ErrNegativeSqrt型は、エラーを呼び出す為の型なので
	// キャストして使用しないと再度Errorが呼び出され、無限ループになる。
	// 実際には実行出来なくなる。
	return fmt.Sprintf("Sqrt(x) is can not ( x < 0 ): %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
