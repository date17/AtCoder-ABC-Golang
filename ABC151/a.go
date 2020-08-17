// 30min
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	alfabet := "abcdefghijklmnopqrstuvwxyz"

	scanner.Scan()
	target := regexp.MustCompile(scanner.Text() + ".")

	answer := target.FindString(alfabet)

	fmt.Println(answer[1:2])

}