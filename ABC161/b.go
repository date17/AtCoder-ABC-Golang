// 25min
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
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	M, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	canSelect := 0
	total := 0
	A := make([]int, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		a, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		A[i] = a

		total += A[i]

	}

	for _, value := range A {
		if value*4*M >= total {
			canSelect++
		}

		if canSelect >= M {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	fmt.Println("No")

}
