// 30min
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main()  {
	// Alfabet 大文字のアルファベットの一巡を定数として定義
	Alfabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	
	scanner := bufio.NewScanner(os.Stdin)

	// 何文字後を示すかを入力
	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	changeAlfabet := make(map[string]string, 26)

	for index, value := range Alfabet {
		targetIndex := index + N
		if targetIndex >= 26 {
			targetIndex -= 26
		}

		changeAlfabet[value] = Alfabet[targetIndex]
	}

	// アルファベット
	scanner.Scan()
	S := strings.Split(scanner.Text(), "")


	answer := ""
	for i := 0; i < len(S); i++ {
		answer += changeAlfabet[S[i]]
	}

	fmt.Println(answer)
}