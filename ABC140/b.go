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
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N-1)

	for i := 0; i < N; i++ {
		scanner.Scan()
		target, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		A[i] = target
	}

	for i := 0; i < N; i++ {
		scanner.Scan()
		target, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		B[i] = target
	}
	for i := 0; i < N-1; i++ {
		scanner.Scan()
		target, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		C[i] = target
	}

	answer := 0
	for i := 0; i < N; i++ {
		answer += B[A[i]-1]

		if i > 0 && A[i]-A[i-1] == 1 {
			answer += C[A[i-1]-1]
		}
	}

	fmt.Println(answer)
}
