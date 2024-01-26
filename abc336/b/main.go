package b

// 必要なパッケージをインポートします
import (
	"fmt"
	"strconv"
)

// main関数を定義します
func main() {
	// 変数nを定義します
	var n int64
	// ユーザーからの入力を受け取ります
	fmt.Scan(&n)

	// nを2進数の文字列に変換します
	s := strconv.FormatInt(n, 2)
	// 文字列sをルーンのスライスに変換します
	r := []rune(s)
	// ルーン"0"を定義します
	t := []rune("0")[0]
	// カウントを初期化します
	count := 0

	// ルーンのスライスを逆順にループします
	for i := len(r) - 1; i >= 0; i-- {
		// ルーンが"0"でない場合、ループを終了します
		if r[i] != t {
			break
		}
		// ルーンが"0"の場合、カウントを増やします
		count++
	}

	// カウントの結果を出力します
	fmt.Println(count)
}
