package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	y, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", x*y)

}
