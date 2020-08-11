// 8min
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	S := scanner.Text()
	scanner.Scan()
	T := scanner.Text()

	scanner.Scan()
	A, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	B, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	U := scanner.Text()
 
	if S == U {
		fmt.Println(A - 1, B)
	} else if T == U {
		fmt.Println(A, B - 1)
	} else {
		fmt.Println(A , B)
	}
}