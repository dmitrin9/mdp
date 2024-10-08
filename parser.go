package main

import (
	"fmt"
	"os"
	"strings"
)

type token_t struct {
	tok_type string
	tok_raw  string
}

type markdownParseNode struct {
	idx1     uint
	idx2     uint
	depth    int
	property string
}

type MarkdownState struct {
	buf            []token_t
	filepath       string
	filedata       [][]byte
	fullfileoffset int
	bufsize        int
	row            int
	col            int

	parseIndex int

	done int

	nodes []markdownParseNode
}

var _MDParser_ExitCodeMap map[int]string = map[int]string{
	100: "E_OK",
	200: "EOF",
	300: "ColumnOverflow",
	400: "MissingFile",
	500: "Invalid File",
}

var _MDParser_Tokens map[string]string = map[string]string{
	"#":  "HEADER",
	"(":  "ENCLOSER",
	")":  "ENCLOSER",
	"[":  "ENCLOSER",
	"]":  "ENCLOSER",
	"=":  "PUNCTUATOR",
	"-":  "PUNCTUATOR",
	"*":  "ENCLOSER",
	"~":  "ENCLOSER",
	"\n": "NEWLINE",
}

func outputError(errno int, filepath string, row int, col int) {
	result := fmt.Sprintf(
		"Parsing error \"%s\" found at %d:%d in %s",
		_MDParser_ExitCodeMap[errno],
		row,
		col,
		filepath,
	)
	fmt.Fprintln(os.Stderr, result)
}

func LoadFile(md *MarkdownState, filepath string) error {
	md.filepath = filepath

	s := strings.Split(md.filepath, ".")
	if s[len(s)-1] != "md" {
		md.done = 500
		outputError(md.done, md.filepath, md.row, md.col)
		return nil
	}

	dat, err := os.ReadFile(md.filepath)
	if err != nil {
		return err
	}

	datArr := []byte(string(dat))
	md.bufsize = len(datArr)

	fileBuffer := [][]byte{}

	firstIdx := 0
	secondIdx := 0
	for i := range string(datArr) {
		if byte(datArr[i]) == byte('\n') {
			secondIdx = i
			fileBuffer = append(fileBuffer, datArr[firstIdx:secondIdx])
			firstIdx = i
		}
	}

	md.filedata = fileBuffer
	md.done = 100
	return nil
}

func next(md *MarkdownState) {
	if md.fullfileoffset >= md.bufsize-1 {
		md.done = 200
		return
	} else if md.col == len(md.filedata[md.row])-1 {
		md.row++
		md.col = 0
	} else {
		md.col++
	}
	md.done = 100
	md.fullfileoffset++
	return
}

func PopulateMarkdownStateBuffer(md *MarkdownState) {
	for _ = range md.bufsize - 1 {
		currentChar := string(md.filedata[md.row][md.col])
		if len(_MDParser_Tokens[currentChar]) > 0 {
			t := token_t{
				tok_type: _MDParser_Tokens[currentChar],
				tok_raw:  currentChar,
			}
			md.buf = append(md.buf, t)
		} else {
			t := token_t{
				tok_type: "LITERAL",
				tok_raw:  currentChar,
			}
			md.buf = append(md.buf, t)
		}
		next(md)
	}
}

func headerParseRule(md *MarkdownState) [][]int {
	// check for no header
	indexBuffer := [][]int{}

	i := md.parseIndex
	for i < len(md.buf) {
		if md.buf[i].tok_raw == "#" && md.buf[i+1].tok_raw != "#" {
			j := i + 1
			for j < len(md.buf) {
				if md.buf[j].tok_type == "NEWLINE" {
					indexBuffer = append(indexBuffer, []int{i + 2, j})
					break
				}
				j++
			}
		} else if md.buf[i].tok_raw == "#" && md.buf[i+1].tok_raw == "#" {
			j := i + 1
			for j < len(md.buf) {
				if md.buf[j].tok_type == "NEWLINE" {
					indexBuffer = append(indexBuffer, []int{i + 3, j})
					break
				}
				j++
			}
		}
		i++
	}
	if i > md.parseIndex {
		md.parseIndex = i
	}
	return indexBuffer
}

func italicParseRule(md *MarkdownState) [][]int {
	indexBuffer := [][]int{}

	i := md.parseIndex
	for i < len(md.buf) {
		if string(md.buf[i].tok_raw) == "*" {
			i += 1
			j := i + 1
			for j < len(md.buf) {
				if string(md.buf[j].tok_raw) == "*" {
					md.buf[i-1] = token_t{tok_raw: " ", tok_type: "LITERAL"}
					md.buf[j] = token_t{tok_raw: " ", tok_type: "LITERAL"}
					indexBuffer = append(indexBuffer, []int{i, j})
					break
				}
				j++
			}
		}
		i++
	}
	if i > md.parseIndex {
		md.parseIndex = i
	}
	return indexBuffer
}

// Iterate until you find a non-thing character and then return a single header, but then if it returns a thing character, make a new one.

/*
func boldParseRule(dat []byte) (int, int) {

}
*/

func ParseMarkdownFromState(md *MarkdownState) {
	/*
		for i := range md.filedata {

		}
	*/
	fmt.Println("Full String ", string(md.filedata[0]))
	val1 := headerParseRule(md)
	val2 := headerParseRule(md)
	val3 := italicParseRule(md)

	fmt.Println("-------------------------------------------")

	//fmt.Println(val1)
	for i := range val1 {
		a := val1[i][0]
		b := val1[i][1]
		fmt.Println("-------------------------")
		fmt.Println(" ", md.buf[a:b])
		fmt.Println("-------------------------")
	}

	fmt.Println("-------------------------------------------")

	//fmt.Println(val2)
	for i := range val2 {
		a := val2[i][0]
		b := val2[i][1]
		fmt.Println("-------------------------")
		fmt.Println(" ", md.buf[a:b])
		fmt.Println("-------------------------")
	}

	fmt.Println("-------------------------------------------")

	//fmt.Println(val3)
	for i := range val3 {
		a := val3[i][0]
		b := val3[i][1]
		fmt.Println("---------------------")
		fmt.Println(md.buf[a:b])
		fmt.Println("---------------------")
	}

}
