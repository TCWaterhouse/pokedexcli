package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "Hello World ALEX"
	actual := cleanInput(words)
	fmt.Print(len(actual))
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
