// 2.5min
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

	fmt.Println(N * N * N)
}
