package main

import "fmt"

func main()  {
	var num uint64
	var num1 uint64

	fmt.Scanln(&num)
	fmt.Scanln(&num1)

	area := (num * num1) / 2

	fmt.Println(area)
}