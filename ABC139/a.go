// 4min
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	S := scanner.Text()

	scanner.Scan()

	T := scanner.Text()

	answer := 0
	for i := 0; i < len(S); i++ {
		if S[i] == T[i] {
			answer++
		}
	}

	fmt.Println(answer)
}
