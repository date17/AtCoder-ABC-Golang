// 10 min
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mul(x int) int {
	ans := 1
	for i := 1; i <= 2; i++ {
		ans *= i
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	M, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// 組み合わせ問題 Nの中から2つのパターンとMの中から2つのパターン
	answer := ((N * (N - 1)) / mul(N)) + ((M * (M - 1)) / mul(M))

	fmt.Println(answer)
}
