package main

import (
	"fmt"
	"fsl-parser/interpreter"
	"fsl-parser/parser"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os"
)

func main() {
	// Get the filenames from command-line arguments
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		log.Fatal("Please provide filenames as command-line arguments.")
	}

	fslInterpreter := interpreter.NewFSLInterpreter()

	for _, file := range filenames {
		// Read the FSL script from a file
		scriptBytes, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading script file: %v\n", err)
			os.Exit(1)
		}

		// Create the ANTLR input stream
		inputStream := antlr.NewInputStream(string(scriptBytes))

		// Create the lexer
		lexer := parser.NewFSLLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the parser
		p := parser.NewFSLParser(tokenStream)

		// Specify custom error handling
		errorListener := antlr.NewDiagnosticErrorListener(true)
		p.RemoveErrorListeners()
		p.AddErrorListener(errorListener)

		// Parse the input script
		tree := p.Program()

		// Visit the parse tree with the fslInterpreter
		fslInterpreter.VisitProgram(tree.(*parser.ProgramContext))
	}

}
