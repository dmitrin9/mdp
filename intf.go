package main

import "fmt"

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

const resetStyleProperty = "\x1B[0m"

func main() {
	fmt.Println(ForegroundColors["magenta"] + BackgroundColors["green"] + italicAnsi + "italic" + resetStyleProperty)
	fmt.Println(boldAnsi + "bold" + resetStyleProperty)
	fmt.Println(faintAnsi + "faint" + resetStyleProperty)
	fmt.Println(underlinedAnsi + "underline" + resetStyleProperty)
	fmt.Println(strikethroughAnsi + "strikethrough" + resetStyleProperty)
	fmt.Println(inverseAnsi + "inverse" + resetStyleProperty)
}
