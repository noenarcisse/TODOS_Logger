package todos

import (
	"TODOS_Logger/pkg/console"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//could go in another slice tbf

// Creates a folder to save log files if it doesnt exist
func createLogDir(dirname string) error {

	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}
	return nil
}

func WriteToConsole(tds []TodoLine) {
	resetColor := "\033[0m"
	sb := strings.Builder{}

	for _, tl := range tds {

		sb.WriteString(string(console.UNDERLINE))
		sb.WriteString(string(console.BLUE))
		sb.WriteString(tl.File)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(tl.LineNum))
		sb.WriteString(resetColor)
		sb.WriteString(string(console.END_U))
		sb.WriteString(" : \n")

		for k, v := range tl.Lines.Items() {

			sb.WriteString(strconv.Itoa(k))
			sb.WriteString(" : ")
			sb.WriteString(string(console.ITALICS))
			sb.WriteString(string(console.GREEN))
			sb.WriteString(v)
			sb.WriteString(resetColor)
			sb.WriteString(string(console.END_I))
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}
func WriteToFile(s string) error {
	t := time.Now()

	logfilename := fmt.Sprintf("log_%d", t.Unix())
	fmt.Printf("Logging in %s\n", logfilename)

	dirname := "logs"
	err := createLogDir(dirname)
	if err != nil {
		return err
	}

	handle, err := os.OpenFile(filepath.Join(dirname, logfilename), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer handle.Close()

	written, err := handle.WriteString(s)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes written\n", written)
	return nil
}
func WriteToSpecialFile(s string, logfilename string, ext string) error {
	filetowrite := logfilename + "." + ext
	fmt.Printf("Logging in %s\n", filetowrite)

	dirname := "logs"
	err := createLogDir(dirname)
	if err != nil {
		return err
	}

	handle, err := os.OpenFile(filepath.Join(dirname, filetowrite), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer handle.Close()

	written, err := handle.WriteString(s)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes written\n", written)
	return nil
}
