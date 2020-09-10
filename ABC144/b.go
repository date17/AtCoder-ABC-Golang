// 5.5min
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

	for i := 1; i < 10; i++ {
		if N/i < 10 && N%i == 0 {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	fmt.Println("No")
}
