package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var a string
	var n int
	var i int
	var err error
	s := make([]int, 0, 50)
	for {
		fmt.Printf("\n\tput an integer to add to slice, or put X to exit:\n\t---> ")
		fmt.Scanln(&a)
		if len(a) == 1 && (a[0] == 'x' || a[0] == 'X') {
			return
		} else {
			n, err = strconv.Atoi(a)
			if err == nil {
				s = append(s, n)
			}
			sort.Ints(s)
			for _, i = range s {
				fmt.Printf("%d ", i)
			}
			fmt.Printf("\n")
		}
	}

}
