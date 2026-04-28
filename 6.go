package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Println(" введи число")
		fmt.Scanln(&x)

		switch {
		case x >= 10:
			println("bid")
		case x >= 1:
			println("medium")
		case x <= 0:
			println("small or zero")
		}
		var value string
		fmt.Println("повторить? yes/no")
		fmt.Scanln(&value)

		if value == "no" {
			break
		}
	}
}
