package main;

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buffer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	wacc := &WACC{Buffer: string(buffer)}
	wacc.Init()

	if err := wacc.Parse(); err != nil {
		log.Fatal(err)
	}
	wacc.PrintSyntaxTree()
}