// 6min
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

	if N%2 == 0 {
		fmt.Println(float32((N / 2)) / float32(N))
	} else {
		fmt.Println(float32(((N + 1) / 2)) / float32(N))
	}
}
