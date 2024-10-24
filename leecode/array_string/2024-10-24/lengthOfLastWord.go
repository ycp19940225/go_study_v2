package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "   fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println(res)
}

func lengthOfLastWord(s string) int {
	var length = 0
	var flags = false
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(string(s[i]) == " ")
		if string(s[i]) != " " && !flags {
			length++
			flags = true
		} else if flags {
			length++
		} else if string(s[i]) == " " && flags {
			break
		}
	}
	s = strings.Trim(s, " ")
	//if length == 0 {
	//	return len(s) - 1
	//}
	return length
}

func lengthOfLastWord2(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
