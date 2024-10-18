package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	md := MarkdownState{}
	LoadFile(&md, file)
	PopulateMarkdownStateBuffer(&md)
	ParseMarkdownFromState(&md)
	fmt.Println(Render(&md))
}
