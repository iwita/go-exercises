package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1

	for _, ch := range input {

		// Solution No. 1
		// min, max := 'A', 'Z'
		// if ch >= min && ch <= max {
		// 	answer++
		// }

		// Solution No. 2
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	fmt.Println(answer)
}
