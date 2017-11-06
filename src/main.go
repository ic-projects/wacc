package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
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

func main() {
	printTree := flag.Bool("t", false, "Prints the AST created by the parser")
	semanticCheck := flag.Bool("s", false, "Parse the file for syntax and symantic errors and generate an AST")
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
