// 6min
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
	answer := ""
	for i := 0; i < len(S); i++ {
		answer += "x"
	}
	fmt.Println(answer)
}
