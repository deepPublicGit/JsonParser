package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		parse(inputFile)

		fmt.Println(inputFile)
	} else {
		fmt.Println("Error: No File Provided!")
	}

}

func parse(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: File may not be present or is corrupt! %s\n", err)
	} else {
		read := bufio.NewReader(file)
		stk := list.New()
		emptyFileFlag := 1
		for {
			if char, _, err := read.ReadRune(); err != nil {
				if err == io.EOF {
					if stk.Len() == 0 && emptyFileFlag == 0 {
						fmt.Println("Json Valid! Exit 0")
					} else {
						fmt.Println("Json Invalid! Exit 1")
					}
					break
				} else {
					fmt.Printf("Error: File could not be parsed, unknown exception occurred! %s\n", err)
				}
			} else {
				//Not an empty file
				emptyFileFlag = 0
				if char == '{' {
					stk.PushBack(char)
				} else if char == '}' {
					stk.Remove(stk.Back())
				}

			}

		}
	}
}
