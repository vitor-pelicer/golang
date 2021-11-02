package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	var hash map[string]string
	hash = make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Printf("Enter your address: ")
	scanner.Scan()
	address = scanner.Text()
	hash["name"] = name
	hash["address"] = address
	jb, _ := json.Marshal(hash)
	fmt.Println(string(jb))

}
