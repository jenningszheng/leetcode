package main

import (
	"fmt"
	goat_latin "leetcode/goat-latin"
)

func main() {
	s := "I speak Goat Latin"
	//for _, x := range s {
	//	fmt.Println(x)
	//}
	ans := goat_latin.ToGoatLatin(s)
	fmt.Println(ans)
}
