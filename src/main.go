package main;

import (
    "fmt"
    "log"
    "os"
    "flag"
)

func main() {
  printTree := flag.Bool("t", false, "Prints the AST created by the parser")
  semanticCheck := flag.Bool("s", false, "Parse the file for syntax and symantic errors and generate an AST")
  flag.Parse()
  var tree interface{}
  if (len(os.Args) > 1) {
    var err error
    tree, err = ParseFile(os.Args[len(os.Args) - 1])
    if err != nil {
        log.Fatal(err)
    }
  } else {
    fmt.Println("Error: No file provided")
  }
  if (*semanticCheck) {

  }
  if (*printTree) {
    fmt.Println(tree)
  }
}
