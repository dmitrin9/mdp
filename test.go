package main

import "fmt"

func testParser() {
	files := []string{"./test.md", "./test2.md"}
	for i := range files {
		md := MarkdownState{}
		LoadFile(&md, files[i])
		PopulateMarkdownStateBuffer(&md)
		ParseMarkdownFromState(&md)
	}
}

func testRender() {
	files := []string{"./test.md", "./test2.md"}
	for i := range files {
		md := MarkdownState{}
		LoadFile(&md, files[i])
		PopulateMarkdownStateBuffer(&md)
		ParseMarkdownFromState(&md)
		fmt.Println(Render(&md))
	}
}

func main() {
	//testParser()
	testRender()
}
