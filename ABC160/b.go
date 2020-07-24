// 5min
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
	X, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	grade500 := (X / 500) * 1000
	grade5 := ((X % 500) / 5) * 5

	fmt.Println(grade500 + grade5)
}
