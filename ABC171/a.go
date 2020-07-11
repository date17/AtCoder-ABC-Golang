package main

import (
	"fmt"
	"regexp"
)

func main() {
	var a string

	fmt.Scan(&a)

	regSmall := regexp.MustCompile(`^[a-z]$`)
	regLarge := regexp.MustCompile(`^[A-Z]$`)

	if regSmall.MatchString(a) {
		fmt.Println("a")
	} else if regLarge.MatchString(a) {
		fmt.Println("A")
	}
}
