package expleto

import (
	// "bytes"
	// "encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	// "strconv"
	// "strings"
	// "text/template"
	// "time"
)

// returns stack trace with source of raised error
func FormatError(err error) error {
	pc, fn, line, _ := runtime.Caller(1)
	return fmt.Errorf("%v\nTRACE: %s[%s:%d]", err, runtime.FuncForPC(pc).Name(), filepath.Base(fn), line)
}

// read data from file
// func GetDataFromFile(path string) ([]byte, error) {
// 	_, err := os.Stat(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// read data from file factory
func getDataFromFileFactory(
	stat func(filename string) (os.FileInfo, error),
	readFile func(filename string) ([]byte, error),
) func(path string) ([]byte, error) {

	return func(path string) ([]byte, error) {
		_, err := stat(path)
		if err != nil {
			return nil, err
		}
		data, err := readFile(path)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

var GetDataFromFile = getDataFromFileFactory(os.Stat, ioutil.ReadFile)
