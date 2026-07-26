package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("input.log")
	if err != nil {
		fmt.Println("ошибка")

	}
	now := time.Now()
	formatted := now.Format("2006-01-02 15:04:05")

	fmt.Fprintf(file, "%s программа запущена\n", formatted)

	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}
		fmt.Fprintln(file, formatted, line)
	}

	fmt.Println(formatted, "готово")

}
