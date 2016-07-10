package expleto

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	// "path/filepath"
	"strings"
	"testing"
	"time"
)

type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
	isdir   bool
	sys     interface{}
}

func New(name string, size int64, mode os.FileMode, modtime time.Time, isdir bool,
	sys interface{}) (fi os.FileInfo) {

	fi = &FileInfo{
		name:    name,
		size:    size,
		mode:    mode,
		modtime: modtime,
		isdir:   isdir,
		sys:     sys,
	}
	return
}

func (fi FileInfo) Name() string {
	return fi.name
}

func (fi FileInfo) Size() int64 {
	return fi.size
}

func (fi FileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi FileInfo) ModTime() time.Time {
	return fi.modtime
}

func (fi FileInfo) IsDir() bool {
	return fi.isdir
}

func (fi FileInfo) Sys() interface{} {
	return fi.sys
}

func TestGetDataFromFile(t *testing.T) {
	// non exists files
	nonexist_cfgFiles := []string{
		"fixtures/nonexist/config/app.json",
		"fixtures/nonexist/config/app.yml",
		"fixtures/nonexist/config/app.toml",
	}
	for _, f := range nonexist_cfgFiles {
		_, err := GetDataFromFile(f)
		if err == nil {
			t.Fatal("There wasn't raise an error")
		}
		str := fmt.Sprintf("%v", err)
		if !strings.Contains(str, "no such file or directory") {
			t.Fatal("The '" + str + "' is not contain the err 'no such file or directory'")
		}
	}

	t_file, err := ioutil.TempFile(os.TempDir(), "temp.json")
	defer os.Remove(t_file.Name())
	if err != nil {
		t.Fatal(err)
	}
	// ioutil.ReadFile = func_name
	_, err = GetDataFromFile(t_file.Name())
	if err != nil {
		t.Fatal("There shouldn't raise an error")
	}

}

func TestMockGetDataFromFile(t *testing.T) {
	// bad
	stat := func(filename string) (os.FileInfo, error) {
		return nil, errors.New("err msg")
	}

	readfile := func(filename string) ([]byte, error) {
		t.Error("should not call this function")
		return nil, nil
	}

	getDataFromFile := getDataFromFileFactory(stat, readfile)

	if _, err := getDataFromFile("foo"); err.Error() != "err msg" {
		t.Error("expected an error to be thrown")
	}

}

func TestMockGetDataFromFileG(t *testing.T) {

	//good
	stat_ok := func(filename string) (os.FileInfo, error) {
		fi := New("file.txt", int64(123), os.ModeType, time.Now(), true, nil)
		return fi, nil
	}
	readfile := func(filename string) ([]byte, error) {

		return nil, errors.New("err msg")
	}
	getDataFromFile := getDataFromFileFactory(stat_ok, readfile)

	if _, err := getDataFromFile("foo"); err.Error() != "err msg" {
		t.Error("expected an error to be thrown")
	}

}
func TestFormatError(t *testing.T) {
	err := errors.New("test 1")
	err1 := FormatError(err)
	err_str := fmt.Sprintf("%v", err)

	err1_str := fmt.Sprintf("%v", err1)
	if err_str == err1_str {
		t.Fatal("err1_str: " + err1_str + " == err_str " + err_str)
	}
}
