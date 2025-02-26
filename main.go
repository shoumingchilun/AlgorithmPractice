package main

import (
	"fmt"
)

func main() {
	for _, c := range "你好" {
		fmt.Println(c)
	}
	fmt.Println(len("你好"))
	fmt.Println("你好"[:3])
	fmt.Println("你好"[3:])

	r := []rune("你好")
	for _, c := range r {
		fmt.Println(c)
	}
}
