package main

import (
	// "fmt"
	"bytes"
	"flag"
	"io"
	// "github.com/eliquious/lexer"
	"github.com/subsilent/crypto/ssh/terminal"
	"os"
)

var debug = flag.Bool("debug", false, "Enable debug logging")

const PROMPT = "\xc4\xa7 >>> "

func main() {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)

	// var scanner *lexer.Scanner
	var inputBuffer bytes.Buffer
	parser := NewParser(&inputBuffer)

	rw := ReadWriter{os.Stdin, os.Stdout}
	term := terminal.NewTerminal(rw, PROMPT)
	resp := ResponseWriter{DefaultColorCodes, term}

	for {
		resp.Write(resp.Colors.Reset)

		line, err := term.ReadLine()
		if err != nil {
			return
		} else if line == "" {
			continue
		} else if line == "quit" || line == "exit" {
			resp.Write(resp.Colors.LightYellow)
			resp.Write([]byte("\n Exiting...\n"))
			resp.Write(resp.Colors.Reset)
			return
		}

		inputBuffer.WriteString(line)
		for {
			expr, err := parser.ParseExpression()
			if err != nil {
				if err == EOL {
					continue
				} else if err == EOF {
					break
				}
				resp.Write(resp.Colors.Red)
				resp.Write([]byte(err.Error() + "\n"))
				resp.Write(resp.Colors.Reset)
			} else if expr != nil {

				// TODO: Evaluate Expression
				result, err := Evaluate(expr)
				if err != nil {
					resp.Write(resp.Colors.Red)
					resp.Write([]byte(err.Error() + "\n"))
					resp.Write(resp.Colors.Reset)
				} else {
					resp.Write(resp.Colors.Green)
					resp.Write([]byte(result + "\n"))
					resp.Write(resp.Colors.Reset)
				}
			}
		}
		inputBuffer.Reset()

		// scanner = lexer.NewScanner(strings.NewReader(line))
		// for {
		// 	tok, pos, lit := scanner.Scan()
		// 	if tok == lexer.EOF || tok == lexer.ILLEGAL {
		// 		break
		// 	} else if tok == lexer.WS {
		// 		continue
		// 	}

		// 	if *debug {
		// 		resp.Write(resp.Colors.LightGrey)
		// 		resp.Write([]byte("DEBUG: "))
		// 		resp.Write(resp.Colors.Green)
		// 		resp.Write([]byte(fmt.Sprintf("%-8s", tok.String())))
		// 		// resp.Write(resp.Colors.White)
		// 		resp.Write(resp.Colors.Yellow)
		// 		resp.Write([]byte(fmt.Sprintf("%3d", pos.Line)))
		// 		resp.Write(resp.Colors.White)
		// 		resp.Write([]byte(":"))
		// 		resp.Write(resp.Colors.Yellow)
		// 		resp.Write([]byte(fmt.Sprintf("%-3d", pos.Char)))
		// 		resp.Write(resp.Colors.White)
		// 		resp.Write([]byte(lit))
		// 		resp.Write([]byte("\n"))
		// 	}
		// }

	}
}

type ReadWriter struct {
	io.Reader
	io.Writer
}
