package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/anurag1109/Golang/01Hello"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")

	input, _ := reader.ReadByte()
	fmt.Println("you entered", string(input))
	fmt.Println("you entered", int('2'))
}
