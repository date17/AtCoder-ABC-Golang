package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		K, A, B                    int
		inputFirst, inputa, inputb string
	)

	fmt.Scan(&inputFirst)
	fmt.Scan(&inputa, inputb)
	K, _ = strconv.Atoi(inputFirst)
	A, _ = strconv.Atoi(inputa)
	B, _ = strconv.Atoi(inputb)

	kMultiple := K * 2

	fmt.Printf("%#v", kMultiple)
	fmt.Printf("%#v", A)
	fmt.Printf("%#v", B)

	// if A < B {
	// 	fmt.Println("OK")
	// } else {
	// 	fmt.Println("NG")
	// }

}
