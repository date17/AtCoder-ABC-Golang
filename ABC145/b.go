// 10min
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

	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()

	S := scanner.Text()

	if N%2 != 0 {
		fmt.Println("No")
		os.Exit(0)
	}
	if S[:N/2] != S[N/2:] {
		fmt.Println("No")
		os.Exit(0)
	}

	fmt.Println("Yes")
}
