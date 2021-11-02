package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Printf("Insert a string: ")
	fmt.Scanf("%s", &s)
	if (strings.HasPrefix(s, "i") || strings.HasPrefix(s, "I")) && (strings.HasSuffix(s, "n") || strings.HasSuffix(s, "N")) && (strings.Contains(s, "a") || strings.Contains(s, "A")) {
		fmt.Printf("Found\n")
	} else {
		fmt.Printf("Not found")
	}
}
