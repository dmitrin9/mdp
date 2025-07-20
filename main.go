package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const INVALID_ARGUMENT = "Invalid value or number of arguments."

func run() error {
	if len(os.Args) < 2 {
		return errors.New(INVALID_ARGUMENT)
	}
	file := os.Args[1]
	md := MarkdownState{}
	LoadFile(&md, file)
	PopulateMarkdownStateBuffer(&md)
	ParseMarkdownFromState(&md)
	fmt.Println(Render(&md))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
