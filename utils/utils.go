package utils

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

func GetDataFromFile(path string) ([]byte, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
