// 11.5min
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

	// N
	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// K
	scanner.Scan()
	K, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// M
	scanner.Scan()
	M, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	line := M * N

	for i := 0; i < N-1; i++ {
		scanner.Scan()
		target, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		line -= target
	}

	if line > K {
		fmt.Println("-1")
	} else if line < 0 {
		fmt.Println("0")
	} else {
		fmt.Println(line)
	}

}
