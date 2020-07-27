// about 3 min
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	S := strings.Split(scanner.Text(), "")

	if S[0] != S[1] || S[1] != S[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
