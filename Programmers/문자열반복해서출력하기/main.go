package main

import "fmt"

func main() {
	var str,emptyString string
	var n int
	fmt.Scan(&str)
	fmt.Scan(&n)
	for i := 0; i < n;i++{
		emptyString += str
	}
	fmt.Println(emptyString)
}