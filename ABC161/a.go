// 4.5min
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func changeAtoB(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func changeAtoC(a *int, c *int) {
	var tmp int
	tmp = *a
	*a = *c
	*c = tmp
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	X, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	Y, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	Z, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	changeAtoB(&X, &Y)
	changeAtoC(&X, &Z)

	fmt.Println(X, Y, Z)

}
