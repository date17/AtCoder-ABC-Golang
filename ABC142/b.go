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
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	K, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	counter := 0

	for i := 0; i < N; i++ {
		scanner.Scan()
		target, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if target >= K {
			counter++
		}
	}

	fmt.Println(counter)
}
