package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//スキャナーの生成
	scanner := bufio.NewScanner(os.Stdin)

	//コマンドからの入力をスキャンする(ついでに整数変換)
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())

	//整数変換のエラー処理
	if err != nil {
		panic(err)
	}

	answer := i + i*i + i*i*i

	fmt.Println(answer)

}

/* bufioの使い方の調査を差し引いて、2min*/
