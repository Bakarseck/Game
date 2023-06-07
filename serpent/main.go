package main

import (
	"fmt"
)

/*
	Rot 13
	- abcdefghijklmnopqrstuvwxyz

	- abcdefghijklmnopqrstuvwxyz
	- tuvwxyzabcdefghijklmnopqrs


	key = 13
*/

func myCypher(str string) string {
	cypher := ""
	for _, v := range str {
		if v >= 97 && v <= 122 {
			cypher += string( v % 26 + 'a' )
			fmt.Println(v%26, v/26)
		} else {
			cypher += string(v)
		}
	}
	return cypher
}

func rot(str string, key byte) string {
	cypher := ""
	for i:=0;i<len(str);i++{
		if str[i] >= 97 && str[i] <= 122 {
			c := str[i] - 'a'
			c = c + key
			c = c % 26
			c = c + 'a'
			cypher += string(c)
		} else if str[i] >= 65 && str[i] <= 90 {
			c := str[i] - 'A'
			c = c + key
			c = c % 26
			c = c + 'A'
			cypher += string(c)
		} else {
			cypher += string(str[i])
		}
	}
	return cypher
}

func main() {
	fmt.Println(rot("abcdefghijklmnopqrstuvwxyz", 13))
}
