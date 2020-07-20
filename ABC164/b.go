// 8min
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func attackTurn(Attack int, HP int) int {
	return HP - Attack
}

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
	C, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	D, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for {
		C = attackTurn(B, C)

		if C <= 0 {
			fmt.Println("Yes")
			break
		}

		A = attackTurn(D, A)
		if A <= 0 {
			fmt.Println("No")
			break
		}
	}
}
