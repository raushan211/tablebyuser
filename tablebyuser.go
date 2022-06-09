package main

import "fmt"

func main() {
	fmt.Println("enter digit")
	var digit int
	fmt.Scanln(&digit)

	printTable(digit)

}
func printTable(digit int) {
	var i int
	j := 1
	count := (digit * 10) + 1
	for i = digit; i < count; i += digit {
		fmt.Println(digit, "*", j, "=", i)
		j = +1

	}
}
