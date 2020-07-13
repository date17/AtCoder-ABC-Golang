package main

import (
	"bufio"
	"os"
)

func main() {
	//標準入力のスキャンの生成
	scanner := bufio.NewScanner(os.Stdin)
	//空白区切りで文字列を取得
	scanner.Split(bufio.ScanWords)

	//N,M,K
	scanner.Scan()

}
