package console

import (
	"fmt"
	"strings"
)

type Color string

const (
	reset  Color = "\033[0m"
	RED    Color = "\033[31m"
	GREEN  Color = "\033[32m"
	YELLOW Color = "\033[33m"
	BLUE   Color = "\033[34m"
)

func Printc(color Color, format string, a ...any) {
	sb := strings.Builder{}
	sb.WriteString(string(color))
	sb.WriteString(format)
	sb.WriteString(string(reset))
	fmt.Printf(sb.String(), a...)
}

func Printcln(color Color, format string, a ...any) {
	sb := strings.Builder{}
	sb.WriteString(string(color))
	sb.WriteString(format)
	sb.WriteString("\n")
	sb.WriteString(string(reset))
	fmt.Printf(sb.String(), a...)
}
