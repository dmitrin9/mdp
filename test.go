package main

func testParser() {
	md := MarkdownState{}
	LoadFile(&md, "./test.md")
	PopulateMarkdownStateBuffer(&md)
	/*
		fmt.Println(md.done)
		fmt.Println(md.buf)
	*/

	ParseMarkdownFromState(&md)
}

func testRender() {
	md := MarkdownState{}
	LoadFile(&md, "./test.md")
	PopulateMarkdownStateBuffer(&md)
	ParseMarkdownFromState(&md)
	Render(&md)
}

func main() {
	testParser()
	testRender()
}
