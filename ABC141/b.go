// many time
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

	for i := 0; i < len(S); i++ {
		if i%2 == 0 {
			if S[i:i+1] == "L" {
				fmt.Println("No")
				os.Exit(0)
			}
		} else {
			if S[i:i+1] == "R" {
				fmt.Println("No")
				os.Exit(0)
			}
		}
	}

	fmt.Println("Yes")
}
