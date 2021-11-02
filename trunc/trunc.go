package main

import (
	"fmt"
)

func main() {
	var Nfloat float32
	fmt.Printf("Insert a float point number: ")
	fmt.Scanf("%f", &Nfloat)
	i := int(Nfloat)
	fmt.Printf("\n%d\n", i)
}
