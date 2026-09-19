package console

import (
	"fmt"
	"strings"
)

// BLACK  => "\033[30m",
// RED    => "\033[31m",
// GREEN  => "\033[32m",
// YELLOW => "\033[33m",
// BLUE   => "\033[34m",
// PURPLE => "\033[35m",
// CYAN   => "\033[36m",
// WHITE  => "\033[37m",

// # background color
// BLACKB  => "\033[40m",
// REDB    => "\033[41m",
// GREENB  => "\033[42m",
// YELLOWB => "\033[43m",
// BLUEB   => "\033[44m",
// PURPLEB => "\033[45m",
// CYANB   => "\033[46m",
// WHITEB  => "\033[47m",

// # bold
// B    => "\033[1m",
// BOFF => "\033[22m",

// # italics
// I => "\033[3m",
// IOFF => "\033[23m",

// # underline
// U => "\033[4m",
// UOFF => "\033[24m",

// # invert
// R => "\033[7m",
// ROFF => "\033[27m",

// # reset
// RESET  => "\033[0m",

type FontStyle string

const (
	ITALICS FontStyle = "\033[3m"
	END_I   FontStyle = "\033[27m"

	UNDERLINE FontStyle = "\033[4m"
	END_U     FontStyle = "\033[24m"
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
