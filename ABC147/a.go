// 4min
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

	limit := 22

	for i := 0; i < 3; i++ {
		scanner.Scan()
		target, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		limit -= target

		if limit <= 0 {
			fmt.Println("bust")
			os.Exit(0)
		}
	}

	fmt.Println("win")
}
