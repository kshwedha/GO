package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func is_true(str string) bool {
	match, _ := regexp.MatchString("[^ab]", str)
	if match {
		fmt.Println("unwanted char found.")
		return false
	}
	if len(str) == 1 {
		return true
	}
	for index, _ := range str {
		if str[index] == 98 {
			if index != len(str)-1 {
				if str[index+1] == 97 {
					return false
				}
			} else {
				return true
			}
		}
	}
	return true
}

func main() {
	fmt.Print("Input the string\n")
	var input string
	// first method
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	x := is_true(input)
	fmt.Print(x)
	// second method
	// reader := bufio.Reader
	// reader.ReadString()
	// third method
	// fmt.Scanln(&input)
}
