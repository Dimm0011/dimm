package main

import "fmt"

func main() {
	var a = 1
	var b = 2
	fmt.Println(a)
	fmt.Println(b)
	//i := 1
	//for i <= 100 {
	for i := 0; i <= 10; i++ {
		a, b = b, a+b
		fmt.Println(a)
		//i++
	}
}
