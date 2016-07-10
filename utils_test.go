package expleto

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	// "path/filepath"
	"strings"
	"testing"
)

func func_name(string) ([]byte, error) {

	return nil, errors.New("Can't read ")

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

func TestFormatError(t *testing.T) {
	err := errors.New("test 1")
	err1 := FormatError(err)
	err_str := fmt.Sprintf("%v", err)

	err1_str := fmt.Sprintf("%v", err1)
	if err_str == err1_str {
		t.Fatal("err1_str: " + err1_str + " == err_str " + err_str)
	}
}
