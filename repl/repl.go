package repl

import (
	"bufio"
	"flink_monkey/lexer"
	"flink_monkey/token"
	"fmt"
	"io"
)

// PROMPT input hint
const PROMPT = ">> "

// Start analyzed the input line and prints out the tokens
func Start(in io.Reader, _ io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Read the line and analyse it
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
