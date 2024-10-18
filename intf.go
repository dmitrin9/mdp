package main

import (
	"fmt"
	"os"
	"strings"
)

const italicAnsi = "\x1B[3m" // rule: *italic*
const boldAnsi = "\x1B[1m"   // rule: **bold**
const underlinedAnsi = "\x1B[4m"
const strikethroughAnsi = "\x1B[9m" // rule: ~~strikethrough~~
const inverseAnsi = "\x1B[7m"

var block string = "█"

var ForegroundColors map[string]string = map[string]string{
	"black":   "\x1B[30m",
	"red":     "\x1B[31m",
	"green":   "\x1B[32m",
	"yellow":  "\x1B[33m",
	"blue":    "\x1B[34m",
	"magenta": "\x1B[35m",
	"cyan":    "\x1B[36m",
	"white":   "\x1B[37m",
}

var BackgroundColors map[string]string = map[string]string{
	"black":   "\x1B[40m",
	"red":     "\x1B[41m",
	"green":   "\x1B[42m",
	"yellow":  "\x1B[43m",
	"blue":    "\x1B[44m",
	"magenta": "\x1B[45m",
	"cyan":    "\x1B[46m",
	"white":   "\x1B[47m",
}

var _MDRender_ExitCodeMap map[int]string = map[int]string{
	100: "E_OK",
	200: "EOF",
	300: "Malformed/Incomplete Input",
}

const resetStyleProperty = "\x1B[0m"

func outputRenderError(errno int, errmsg string) {
	result := fmt.Sprintf(
		"Render error: %d\nSuccinct: \"%s\"\nVerbose: \"%s\"",
		errno,
		_MDRender_ExitCodeMap[errno],
	)
	fmt.Fprintln(os.Stderr, result)
}

func sliceBufferToString(md *MarkdownState, i1 int, i2 int) string {
	buffer := []string{}
	for i1 < i2 {
		buffer = append(buffer, md.buf[i1].tok_raw)
		i1++
	}
	return strings.Join(buffer[:], "")
}

func formatItalic(txt string) string {
	return italicAnsi + txt + resetStyleProperty
}

func formatInverse(txt string) string {
	return inverseAnsi + txt + resetStyleProperty
}

func Render(md *MarkdownState) string {
	stringBuilder := []string{}

	cursor := 0
	for i := range md.nodes {
		if md.nodes[i].property == "ITALIC" {
			start, end := md.nodes[i].idx1, md.nodes[i].idx2
			before := sliceBufferToString(md, cursor, start)
			inside := sliceBufferToString(md, start+1, end)
			cursor = end + 1
			stringBuilder = append(stringBuilder, before)
			stringBuilder = append(stringBuilder, formatItalic(inside))
			fmt.Println("BEFORE ", before)
			fmt.Println("INSIDE ", inside)
		} else if md.nodes[i].property == "HEADING" {
			start, end := md.nodes[i].idx1, md.nodes[i].idx2
			before := sliceBufferToString(md, cursor+1, start-1)
			inside := sliceBufferToString(md, start, end+1)
			cursor = end + 1
			stringBuilder = append(stringBuilder, before+formatInverse(inside))
		}
	}
	stringBuilder = append(stringBuilder, sliceBufferToString(md, cursor, len(md.buf)))

	dat := strings.Join(stringBuilder[:], "")
	return dat
}
