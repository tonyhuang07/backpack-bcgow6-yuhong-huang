package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting...")
	file, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	} else {
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(line)
		}
	}
	fmt.Println("ending...")
}
