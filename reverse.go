package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var lines []string

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}
}
