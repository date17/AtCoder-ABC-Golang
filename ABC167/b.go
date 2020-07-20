// 6min
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	A, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	scanner.Scan()
	B, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	scanner.Scan()
	_, err = strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	scanner.Scan()
	K, err := strconv.Atoi(scanner.Text())

	if err != nil {
		panic(err)
	}

	if A >= K {
		fmt.Println(K)
	} else {
		fmt.Println(A + (K-A-B)*(-1))
	}
}
