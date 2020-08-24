// 7min
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	A, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	B, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	K, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	if K >= (A + B) {
		fmt.Println(0, 0)
	} else if A >= K {
		fmt.Println(A-K, B)
	} else {
		B -= (K - A)
		fmt.Println(0, B)
	}

}
