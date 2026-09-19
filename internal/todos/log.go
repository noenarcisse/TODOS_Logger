package todos

import (
	"html"
	"strings"
)

// todo wip
type FormatFunc func(TodoLine) string

func createLog(tds []TodoLine, f FormatFunc) string {
	sb := strings.Builder{}
	for _, td := range tds {
		sb.WriteString(f(td))
		sb.WriteString("\n")
	}
	return sb.String()
}

// Create a full text log based on the TODOS found
func CreateLog(tds []TodoLine) string {
	return createLog(tds, FormatToConsole)
}
func CreateLogToMd(tds []TodoLine) string {
	return createLog(tds, FormatToMd)
}
func CreateLogToHTML(tds []TodoLine) string {
	return createLog(tds, FormatToHTML)
}

func PrepareHTMLContent(files []string, templateHtml string, css string, log string, logfilename string) string {
	sb := strings.Builder{}
	for _, f := range files {
		sb.WriteString("<li>")
		sb.WriteString(html.EscapeString(f))
		sb.WriteString("</li>")
	}
	html2 := strings.Replace(templateHtml, "[css]", css, 1)
	html2 = strings.Replace(html2, "[log]", log, 1)
	html2 = strings.ReplaceAll(html2, "[filename]", strings.ToUpper(logfilename))
	html2 = strings.Replace(html2, "[files]", sb.String(), 1)
	html2 = strings.Replace(html2, "[log]", log, 1)
	return html2
}
