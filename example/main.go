package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jacalz/eval"
)

func main() {
	fmt.Print("Enter mathematical expression: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	result, err := eval.Evaluate(line)
	if err != nil {
		panic(err)
	}

	fmt.Println("The result is:", result)
}
