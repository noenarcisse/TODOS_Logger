package todos

import (
	"html"

	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func FormatToConsole(tl TodoLine) string {
	sb := strings.Builder{}

	sb.WriteString(tl.File)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(tl.LineNum))
	sb.WriteString(" : \n")

	for k, v := range tl.Lines.Items() {
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString(" : ")
		sb.WriteString(v)
		sb.WriteString("\n")
	}
	return sb.String()
}

func FormatToFile(tl TodoLine) string {
	sb := strings.Builder{}

	sb.WriteString(tl.File)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(tl.LineNum))
	sb.WriteRune('\n')

	for k, v := range tl.Lines.Items() {
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString(" : ")
		sb.WriteString(v)
		sb.WriteString("\n")
	}
	return sb.String()
}

func FormatToMd(tl TodoLine) string {
	sb := strings.Builder{}

	filename := fmt.Sprintf("[%s:%s](vscode://file/%s:%s)", tl.File, strconv.Itoa(tl.LineNum), tl.File, strconv.Itoa(tl.LineNum))

	sb.WriteString(filepath.ToSlash(filename))
	sb.WriteString(" <br> \n")
	sb.WriteString("```\n")

	for k, v := range tl.Lines.Items() {
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString(" : ")
		sb.WriteString(v)
		sb.WriteString("\n")
	}
	sb.WriteString("```")
	sb.WriteString("\n<br>\n")
	return sb.String()
}

func FormatToHTML(tl TodoLine) string {
	sb := strings.Builder{}

	//vscode://
	filename := fmt.Sprintf("<a href=\"vscode://file/%s:%s\">%s:%s</a>",
		html.EscapeString(tl.File),
		strconv.Itoa(tl.LineNum),
		html.EscapeString(tl.File),
		strconv.Itoa(tl.LineNum))
	sb.WriteString("<div>\n")
	sb.WriteString("📄 ")
	sb.WriteString(filename)
	sb.WriteString("<br/>\n")
	sb.WriteString("<code>")

	for k, v := range tl.Lines.Items() {

		sb.WriteString("<br/>\n")
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString("   ")
		sb.WriteString(html.EscapeString(v))
	}
	sb.WriteString("</code>")
	sb.WriteString("</div>\n")
	return sb.String()
}
