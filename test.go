package main

import (
	"fmt"
)

func testParser() {
	md := MarkdownState{}
	LoadFile(&md, "./test.md")
	PopulateMarkdownStateBuffer(&md)
	fmt.Println(md.done)
	fmt.Println(md.buf)

	ParseMarkdownFromState(&md)
}

func main() {
	testParser()
}
