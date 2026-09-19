package todos

import (
	"fmt"
	"os"
	"path/filepath"
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

func WriteToConsole(s string) {
	fmt.Println(s)
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
