package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanear := bufio.NewScanner(os.Stdin)

	var x1, x2, x3, x4, x5 int
	var err error

	scanear.Split(bufio.ScanWords)

	//標準入力を取得
	scanear.Scan()
	x1, err = strconv.Atoi(scanear.Text())
	if err != nil {
		panic(err)
	}

	scanear.Scan()
	x2, err = strconv.Atoi(scanear.Text())
	if err != nil {
		panic(err)
	}

	scanear.Scan()
	x3, err = strconv.Atoi(scanear.Text())
	if err != nil {
		panic(err)
	}

	scanear.Scan()
	x4, err = strconv.Atoi(scanear.Text())
	if err != nil {
		panic(err)
	}

	scanear.Scan()
	x5, err = strconv.Atoi(scanear.Text())
	if err != nil {
		panic(err)
	}

	if x1 == 0 {
		fmt.Println(1)
	} else if x2 == 0 {
		fmt.Println(2)
	} else if x3 == 0 {
		fmt.Println(3)
	} else if x4 == 0 {
		fmt.Println(4)
	} else if x5 == 0 {
		fmt.Println(5)
	}

	os.Exit(0)
}
