package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")

	input, _ := reader.ReadByte()
	fmt.Println("you entered", string(input))
	fmt.Println("you entered", int('2'))
}
