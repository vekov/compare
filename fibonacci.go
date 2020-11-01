package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	num := 40
	for i := 0; i < num; i++ {
		fmt.Println(fib(i))
	}
}
