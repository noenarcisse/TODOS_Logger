package todos

import (
	sm "TODOS_Logger/pkg/sorted_map"
	"bufio"
	"os"
	"path/filepath"

	"strings"
	"unicode"
)

//opens file, find lines, caches them

func getCommentSymbol(ext string) string {
	comments := map[string]string{
		".sql": "--",
		".py":  "#",
		".nim": "#",
	}

	commentSymbol := "//" //default val
	if _, ok := comments[ext]; ok {
		commentSymbol = comments[ext]
	}

	return commentSymbol
}

// Checks if a line has a todo comment in it
func HasLineTodo(ext string, s string) bool {
	commentSymbol := getCommentSymbol(ext)
	todo := []rune(commentSymbol + "todo")
	index := 0
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}

		if unicode.ToLower(c) == todo[index] {
			index++

			if index == len(todo) {
				return true
			}
		} else {
			index = 0
		}
	}
	return false
}

func IsCommentLine(ext string, line string) bool {

	commentSymbol := getCommentSymbol(ext)
	commentRunes := []rune(commentSymbol)
	index := 0
	runes := []rune(line)

	for _, c := range runes {
		if unicode.IsSpace(c) {
			continue
		}
		if c == commentRunes[index] {
			index++
			if index == len(commentRunes) {
				return true
			}
		}
	}
	return false
}

// Opens a file, look for TODOS comments and theirs trailing comments.
// Returns all lines founds and closes the file.
func GetAll(file string) (found []TodoLine) {
	handle, err := os.OpenFile(file, os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	scan := bufio.NewScanner(handle)
	scan.Split(bufio.ScanLines)

	linenum := 1
	todo := TodoLine{
		File:  file,
		Lines: sm.New[int, string](),
	}

	isTrailingLine := false
	ext := filepath.Ext(file)

	for scan.Scan() {

		line := scan.Text()

		if isTrailingLine {
			if IsCommentLine(ext, line) {
				todo.Lines.Add(linenum, strings.TrimSpace(line))
			} else {
				found = append(found, todo)
				isTrailingLine = false
				//reset todo
				todo = TodoLine{
					File:  file,
					Lines: sm.New[int, string](),
				}
			}
		}

		if HasLineTodo(ext, line) {
			todo.Lines.Add(linenum, strings.TrimSpace(line))
			todo.LineNum = linenum
			isTrailingLine = true
		}

		if scan.Err() != nil {
			panic(scan.Err())
		}
		linenum++
	}
	return
}
