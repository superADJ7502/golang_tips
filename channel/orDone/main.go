package main

import "fmt"

func main() {
	var val interface{}
	val = <-Or()
	fmt.Println(val)
}
