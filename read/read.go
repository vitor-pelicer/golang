package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type words struct {
	fname string
	lname string
}

func main() {
	var name string
	var line string
	var wl words
	var wslc []words
	wslc = make([]words, 0, 100)
	fmt.Printf("Enter the file name: ")
	fmt.Scanf("%s", &name)
	file, err := os.Open(name)
	if err == nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line = scanner.Text()
			allwords := strings.Split(line, " ")
			wl.fname = allwords[0]
			wl.lname = allwords[len(allwords)-1]
			wslc = append(wslc, wl)
		}
		file.Close()
		for _, w := range wslc {
			fmt.Println(w)
		}

	} else {
		fmt.Printf("error opening file")
	}

}
