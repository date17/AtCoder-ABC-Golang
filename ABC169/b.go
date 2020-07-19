package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	answer := big.NewInt(1)
	arg := big.NewInt(int64(math.Pow10(18)))

	// N
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		num := new(big.Int)
		num, err := num.SetString(scanner.Text(), 10)
		if err == false {
			panic(err)
		}
		if num.Cmp(big.NewInt(int64(0))) == 0 {
			answer.Mul(answer, num)
			break
		}
		answer.Mul(answer, num)
	}

	if answer.Cmp(arg) == 1 {
		fmt.Println("-1")
	} else {
		fmt.Println(answer)
	}

}
