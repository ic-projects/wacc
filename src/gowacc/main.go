package main

import (
	"ast"
	"flag"
	"fmt"
	"grammar"
	"os"
	"path"
	"strings"
)

func main() {
	// Parse arguments
	parseOnly := flag.Bool(
		"parse",
		false,
		"Parse the file for syntax and symantic errors and generate an AST.",
	)
	printTree := flag.Bool(
		"ast",
		false,
		"Display AST generated by the parser.",
	)
	symbolTable := flag.Bool(
		"table",
		false,
		"Displays the symbol table generated by the semantic analysis.",
	)
	debugMode := flag.Bool(
		"debug",
		false,
		"Print out additional debugging information during compilation.",
	)
	semanticOnly := flag.Bool(
		"semantic",
		false,
		"Parse the file for syntax and semantic errors and generate an AST.",
	)
	printAssembly := flag.Bool(
		"asm",
		false,
		"Display ARM assembly code generated by the code generator.",
	)
	filepath := os.Args[len(os.Args)-1]
	flag.Parse()
	if *debugMode {
		ast.DebugMode = true
	}

	// Load the file and parse into an AST
	var tree ast.ProgramNode
	if len(os.Args) > 1 {
		fmt.Println("-- Compiling...")
		var err error
		treeValue, err := grammar.ParseFile(filepath)

		if err != nil {
			fmt.Println("Errors detected during compilation! " +
				"Exit code 100 returned.")
			fmt.Println(err)
			os.Exit(100)
		}

		tree = treeValue.(ast.ProgramNode)

	} else {
		fmt.Println("Error: No file provided")
	}

	if !*parseOnly {

		// Perform semantic error checking
		checker := PerformSemanticCheck(tree)

		// Print out all semantic errors that occur
		if checker.hasErrors() || SimplifyTree(tree, checker).hasErrors() {
			fmt.Println("Errors detected during compilation! Exit code 200 returned.")
			checker.PrintErrors(filepath)
			os.Exit(200)
		}

		// Print out the final symbol table
		if *symbolTable {
			checker.PrintSymbolTable()
		}

		if !*semanticOnly {

			// Generate assembly Code
			asm := GenerateCode(tree, checker.SymbolTable())

			// Save assembly code to files
			savepath := strings.TrimSuffix(path.Base(filepath), ".wacc") + ".s"
			err := asm.SaveToFile(savepath)
			if err != nil {
				fmt.Println("Error creating output file.")
				fmt.Println(err)
				os.Exit(200)
			}

			// Print assembly code
			if *printAssembly {
				fmt.Println("-- Printing Assembly...")
				fmt.Print(savepath)
				fmt.Println(" contents are:")
				fmt.Println("================================================" +
					"===========")
				fmt.Print(asm.NumberedCode())
				fmt.Println("================================================" +
					"===========")
			}
		}
	}

	// Print the AST
	if *printTree {
		fmt.Println("-- Printing AST...")
		fmt.Print(strings.TrimSuffix(path.Base(filepath), ".wacc"))
		fmt.Println(".ast contents are:")
		fmt.Println("========================================================" +
			"===")
		fmt.Print(tree)
		fmt.Println("========================================================" +
			"===")
	}
}
