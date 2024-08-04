package main

import (
	"fmt"
	"os"

	"github.com/danbrakeley/hai/internal/repl"
)

func main() {
	fmt.Println("This is the Hai programming language!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
