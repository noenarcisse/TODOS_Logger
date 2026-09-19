package main

import (
	ae "TODOS_Logger/internal/app_error"
	"TODOS_Logger/internal/help"
	"TODOS_Logger/internal/todos"
	"TODOS_Logger/internal/walker"
	"TODOS_Logger/pkg/console"
	_ "embed"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"time"
)

//go:embed style.css
var css string

//go:embed template.html
var templateHtml string

type set[T comparable] = map[T]struct{}

func main() {

	defer func(s time.Time) {
		fmt.Println("Exec time:", time.Since(s))
	}(time.Now())

	args := os.Args[1:]
	if len(args) <= 0 {
		console.Printcln(console.RED, ae.ErrMsg(ae.MISSING_ARGS))
		help.Display()
		return
	}

	if args[0] == "-h" {
		help.Display()
		return
	}

	splet := strings.Split(args[0], ",")

	if len(splet) <= 0 {
		console.Printcln(console.RED, ae.ErrMsg(ae.NO_FILE_SEARCH_SPECIFIED))
		return
	}

	exts := set[string]{}
	sb := strings.Builder{}
	for _, s := range splet {
		sb.WriteRune('.')
		sb.WriteString(s)
		exts[sb.String()] = struct{}{}
		sb.Reset()
	}

	fmt.Print("Looking for files : ")
	fmt.Println(strings.Join(slices.Collect(maps.Keys(exts)), ", "))

	ignores := set[string]{
		".git":         struct{}{},
		".vscode":      struct{}{},
		"bin":          struct{}{},
		"obj":          struct{}{},
		"node_modules": struct{}{},
		".venv":        struct{}{},
		"__pycache__":  struct{}{},
	}

	files, err := walker.WalkThisWay(".", exts, ignores)
	if err != nil {
		panic(err)
	}

	if len(files) <= 0 {
		console.Printcln(console.RED, ae.ErrMsg(ae.NO_FILE_FOUND))
		return
	}

	todoLines := []todos.TodoLine{}

	for _, f := range files {
		fmt.Println(f)

		found := todos.GetAll(f)
		if len(found) > 0 {
			todoLines = append(todoLines, found...)
		}
	}

	if len(todoLines) <= 0 {
		console.Printcln(console.RED, ae.ErrMsg(ae.NO_TODO_COMMENT))
		return
	}

	//todo refacto mieux ca en lazy ou equivalent ou prep un bloc d'ecriture separé de la console
	t := time.Now()
	logfilename := fmt.Sprintf("log_%d", t.Unix())

	if len(args) > 1 {
		options := set[string]{}
		for _, e := range args[1:] {
			options[e] = struct{}{}
		}

		if _, ok := options["c"]; ok {
			// log := todos.CreateLog(todoLines) //todo maybe be unused for md file or html
			fmt.Println("----------------")
			todos.WriteToConsole(todoLines)
			delete(options, "c")
		}

		if _, ok := options["f"]; ok {
			log := todos.CreateLog(todoLines) //todo maybe be unused for md file or html
			err := todos.WriteToFile(log)
			if err != nil {
				panic(err)
			}
			delete(options, "f")
		}

		if _, ok := options["md"]; ok {
			log := todos.CreateLogToMd(todoLines)
			err := todos.WriteToSpecialFile(log, logfilename, "md")
			if err != nil {
				panic(err)
			}
			delete(options, "md")
		}
		if _, ok := options["html"]; ok {

			log := todos.CreateLogToHTML(todoLines)
			html2 := todos.PrepareHTMLContent(files, templateHtml, css, log, logfilename)
			todos.WriteToSpecialFile(html2, logfilename, "html")
			delete(options, "html")
		}

		if len(options) > 0 {
			console.Printcln(console.RED, "Erreur args non reconnus : %v",
				strings.Join(slices.Collect(maps.Keys(options)), ", "))
		}

	} else {
		// log := todos.CreateLog(todoLines) //todo maybe be unused for md file or html

		todos.WriteToConsole(todoLines)
	}
}
