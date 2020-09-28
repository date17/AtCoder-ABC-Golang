// 14min
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

	answerMax := 0
	counter := 0

	scanner.Scan()
	target, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < N; i++ {
		scanner.Scan()
		nextTarget, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if target-nextTarget >= 0 {
			counter++
			if answerMax < counter {
				answerMax = counter
			}
		} else {
			if answerMax < counter {
				answerMax = counter
			}
			counter = 0
		}
		target = nextTarget
	}

	fmt.Println(answerMax)
}
