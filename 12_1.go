package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)
	words := strings.Fields(text)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"слово", "частота"})

	for word, count := range wordCount {
		writer.Write([]string{word, strconv.Itoa(count)})
	}
	writer.Flush()
}
