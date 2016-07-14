package expleto

// This file is part of Expleto, a web content management system.
// Copyright 2016 Valeriy Solovyov <weldpua2008@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an "AS IS"
// BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
