// 3min
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

	r, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r * r)
}
