package main

import (
	"os"

	"./Repl"
)

func main() {
	Repl.Start(os.Stdin, os.Stdout)
}
