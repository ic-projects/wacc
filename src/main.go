package main;

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tree, err := ParseFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tree)
}