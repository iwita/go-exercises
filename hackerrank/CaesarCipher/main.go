package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// Solution No. 1
	// newRune := rotate('z', 2, alphabet)
	// fmt.Println(string(newRune))

	var ret []rune
	for _, ch := range input {
		ret = append(ret, rotateAny(ch, delta))
	}
	fmt.Println(string(ret))

}

func rotateAny(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotateWithBase(r, delta, 'A')
	} else if r >= 'a' && r <= 'z' {
		return rotateWithBase(r, delta, 'a')
	} else {
		return r
	}
}

func rotateWithBase(r rune, delta int, base int) rune {

	tmp := int(r) - base
	tmp = (tmp + delta) % 26

	return rune(tmp + base)

}

func rotate(s rune, delta int, key []rune) rune {

	idx := strings.IndexRune(string(key), s)
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }

	idx = (idx + delta) % len(key)

	// if idx < 0 {
	// 	panic("idx < 0")
	// }
	// for i := 0; i < delta; i++ {
	// 	idx++
	// 	if idx > len(key) {
	// 		idx = 0
	// 	}
	// }
	return key[idx]
}
