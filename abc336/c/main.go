package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	// bufio.NewReaderとbufio.NewWriterを使用して、標準入力と標準出力を効率的に扱うための準備を行います。
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	// プログラムの終了時に、バッファに溜まった出力をフラッシュ（実際に出力）します。
	defer out.Flush()

	// 整数Nを標準入力から読み込みます。
	var n int
	fmt.Fscan(in, &n)

	// 5の累乗の数を格納するマップを作成
	// 5進法での各桁の値を格納するためのマップを作成します。
	digits := make(map[int]int, 18)
	for i := 17; i >= 0; i-- {
		// 5のi乗を計算
		// 5進法での各桁が表す値を計算します。
		digit := int(math.Pow(5, float64(i)))
		if n >= digit {
			if n < 5 {
				// n-1を格納
				// nが5未満の場合、その桁に配置するべき数字はn-1となります。
				digits[i] = n - 1
				break
			}

			// nをdigitで割った商（count）を計算します。
			count := n / digit
			if n%digit == 0 {
				// nがdigitで割り切れる場合、count-1を格納し、以降の桁を4で埋める
				// その桁に配置するべき数字はcount-1となります。
				// また、以降の桁はすべて4（5進法での最大値）で埋めます。
				digits[i] = count - 1
				for s := i - 1; s >= 0; s-- {
					digits[s] = 4
				}
				break
			}
			// nがdigitで割り切れない場合、countを格納し、nからdigit*countを引く
			// その桁に配置するべき数字はcountとなります。
			// また、nからdigit*countを引いて、次の桁の計算に備えます。
			digits[i] = count
			n -= digit * count
			continue
		}

		// nがdigit未満の場合、0を格納
		// その桁に配置するべき数字は0となります。
		digits[i] = 0
	}

	// 5進法での各桁の値を10進法での偶数（0, 2, 4, 6, 8）にマッピングして、求めるべき「良い整数」を計算します。
	ans := 0
	numbers := [5]int{0, 2, 4, 6, 8}
	for index, value := range digits {
		if index == 0 {
			// indexが0の場合、numbers[value]を直接加算します。
			ans += numbers[value]
		} else {
			// indexが0以外の場合、numbers[value] * 10のindex乗を加算します。
			ans += numbers[value] * int(math.Pow10(index))
		}
	}

	// 計算した「良い整数」を標準出力に出力します。
	fmt.Fprintln(out, ans)
}

// 2つの整数aとbを引数に取り、その最大値を返す関数です。
func max(a, b int) int {
	if a > b {
		// aがbより大きい場合、aを返します。
		return a
	}
	// aがb以下の場合、bを返します。
	return b
}
