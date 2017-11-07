package main

import (
	"ast"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"bufio"
)

func number(s string) string {
	var buf bytes.Buffer
	for i, line := range strings.Split(s, "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%d\t%s\n", i, line))
		}
	}
	return buf.String()
}

func getLine(path string, n int) string {
	// Open the WACC source file
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: Unable to open the specified WACC source file")
		os.Exit(100)
	}

	reader := bufio.NewReader(f)
	var line string

	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Unable to read specified line")
			os.Exit(100)
		}
	}

	return line
}

func main() {
	printTree := flag.Bool("t", false, "Prints the AST created by the parser")
	semanticCheck := flag.Bool("s", true, "Parse the file for syntax and symantic errors and generate an AST")
	filepath := os.Args[len(os.Args)-1]
	flag.Parse()
	var tree interface{}
	if len(os.Args) > 1 {
		fmt.Println("-- Compiling...")
		var err error
		tree, err = ParseFile(filepath)
		if err != nil {
			fmt.Println("Errors detected during compilation! Exit code 100 returned.")
			fmt.Println(err)
			os.Exit(100)
		}
	} else {
		fmt.Println("Error: No file provided")
	}
	if *semanticCheck {
		checker := ast.NewSemanticCheck()
		ast.Walk(checker, tree)

		// Print out all errors that occur
		if len(checker.Errors) > 0 {
			for i, e := range checker.Errors {
				if (i > 5) {
					fmt.Printf("And %d other errors\n", len(checker.Errors) - 5)
					break
				}

				var b bytes.Buffer
				b.WriteString("\nSemantic Error at ")
				b.WriteString(fmt.Sprintf("%s\n", e.Pos))

				b.WriteString(getLine(filepath, e.Pos.LineNumber()))
				b.WriteString(strings.Repeat(" ", e.Pos.ColNumber() - 1))
				b.WriteString("^\n")

				b.WriteString("Expected type ")
				i := 1
				for t := range e.Expected {
					if i == len(e.Expected) {
						b.WriteString(fmt.Sprintf("\"%s\"", t))
					} else {
						b.WriteString(fmt.Sprintf("\"%s\" or ", t))
					}
					i++
				}

				b.WriteString(fmt.Sprintf(" but got type \"%s\"", e.Got))
				fmt.Println(b.String())
			}
			os.Exit(200)
		}
	}
	if *printTree {
		fmt.Println("-- Printing AST...")
		fmt.Print(strings.TrimSuffix(path.Base(filepath), ".wacc"))
		fmt.Println(".ast contents are:")
		fmt.Println("===========================================================")
		fmt.Print(number(fmt.Sprintf("%s", tree)))
		fmt.Println("===========================================================")
	}
	fmt.Println("-- Finished")
}
