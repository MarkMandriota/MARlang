package Repl

import (
	"bufio"
	"fmt"
	"io"

	. "../Lexer"
)

const (
	PROMPT = ">> "
	LOGO   = `
 ██▒      ██▒   ██▒             ████▒
█▒█▒  █▒█▒  █▒█▒          █▒   ██▒
█▒ █▒█▒ █▒  █▒   █▒       ████▒
█▒  ██▒   █▒  █████▒    █▒  █▒
█▒              █▒  █▒        █▒  █▒     █▒
`
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Printf(LOGO)

	for {
		fmt.Fprintf(out, PROMPT)

		scanner.Scan()
		str := scanner.Text()

		for _, el := range New([]byte(str)).Lex() {
			fmt.Fprintf(out, "%d: %s\n", el.Type, el.Val)
		}
	}
}
