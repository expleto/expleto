package expleto

import (
	// "errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	var fixtures = struct {
		AppName   string
		BaseURL   string
		Port      int
		Verbose   bool
		StaticDir string
		ViewsDir  string
	}{
		"expleto web app", "http://localhost:9000", 9000, false, "static", "views",
	}

	if cfg.AppName != fixtures.AppName {
		t.Fatal("cfg.AppName != fixtures.AppName")
	}
	if cfg.BaseURL != fixtures.BaseURL {
		t.Fatal("cfg.BaseURL != fixtures.BaseURL")
	}
	if cfg.Port != fixtures.Port {
		t.Fatal("cfg.Port != fixtures.Port")
	}
	if cfg.Verbose != fixtures.Verbose {
		t.Fatal("cfg.Verbose != fixtures.Verbose")
	}
	if cfg.StaticDir != fixtures.StaticDir {
		t.Fatal("cfg.StaticDir != fixtures.StaticDir")
	}
	if cfg.ViewsDir != fixtures.ViewsDir {
		t.Fatal("cfg.ViewsDir != fixtures.ViewsDir")
	}
}

func TestConfig(t *testing.T) {
	// right files
	// cfgFiles := []string{
	// 	"fixtures/config/app.json",
	// 	"fixtures/config/app.yml",
	// 	"fixtures/config/app.toml",
	// }
	good_store_prefix := "./fixtures/config/good"

	cfgFiles, _ := ioutil.ReadDir(good_store_prefix)
	if len(cfgFiles) < 1 {
		t.Fatalf("Failed because you should have a test cases")
	}
	cfg := DefaultConfig()
	for _, f := range cfgFiles {
		file_path, _ := filepath.Abs(good_store_prefix + "/" + f.Name())
		// t.Log(file_path)
		nCfg, err := NewConfig(file_path)
		if err != nil {
			t.Fatal(err)
		}
		if nCfg.AppName != cfg.AppName {
			t.Errorf("expected %s got %s", cfg.AppName, nCfg.AppName)
		}
	}
	// non-exist files
	nonexist_cfgFiles := []string{
		"fixtures/nonexist/config/app.json",
		"fixtures/nonexist/config/app.yml",
		"fixtures/nonexist/config/app.toml",
	}
	for _, f := range nonexist_cfgFiles {
		_, err := NewConfig(f)
		if err == nil {
			t.Fatal("There wasn't raise an error")
		}

	}
	// wrong syntax
	bad_store_prefix := "./fixtures/config/bad"
	bad_cfgFiles, _ := ioutil.ReadDir(bad_store_prefix)
	if len(bad_cfgFiles) < 1 {
		t.Fatalf("Failed because you should have a test cases")
	}
	for _, f := range bad_cfgFiles {
		file_path, _ := filepath.Abs(bad_store_prefix + "/" + f.Name())
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			t.Fatal("Can't find the bad config file " + file_path)
		}
		_, err := NewConfig(file_path)
		if err == nil {
			t.Fatal("There wasn't raise an error")
		}

	}

	// wrong format
	wrong_store_prefix := "./fixtures/config/wrong_format"
	wrong_cfgFiles, _ := ioutil.ReadDir(wrong_store_prefix)
	if len(wrong_cfgFiles) < 1 {
		t.Fatalf("Failed because you should have a test cases")
	}
	for _, f := range wrong_cfgFiles {
		file_path, _ := filepath.Abs(wrong_store_prefix + "/" + f.Name())
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			t.Fatal("Can't find the wrong config file " + file_path)
		}
		_, err := NewConfig(file_path)
		if err == nil {
			t.Fatal("There wasn't raise an error")
		}
		if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", fmt.Errorf("Can't parse %s: %v", file_path, ERROR_FORMAT_NOT_SUPPORTED)) {
			t.Fatal("Raised unexpected error %v", err)
		}

	}
}

func TestConfigEnv(t *testing.T) {
	os.Clearenv()
	fields := []struct {
		name, env, value string
	}{
		{"AppName", "APP_NAME", "expleto"},
		{"BaseURL", "BASE_URL", "http://localhost:9000"},
		{"Port", "PORT", "9009"},
		{"ViewsDir", "VIEWS_DIR", "viewTest"},
		{"StaticDir", "STATIC_DIR", "statics"},
		{"Verbose", "VERBOSE", "true"},
	}
	for _, f := range fields {

		// check out env name maker
		cm := getEnvName(f.name)
		if cm != f.env {
			t.Errorf("expected %s got %s", f.env, cm)
		}
	}

	// set environment values
	for _, f := range fields {
		_ = os.Setenv(f.env, f.value)
	}

	cfg := DefaultConfig()
	if err := cfg.Sync(); err != nil {
		t.Errorf("Can't syncing env %v", err)
	}

	if cfg.Port != 9009 {
		t.Errorf("expected 9000 got %d instead", cfg.Port)
	}

	if cfg.Verbose != true {
		t.Errorf("expected expleto got %s", cfg.Verbose)
	}
	if cfg.AppName != "expleto" {
		t.Errorf("expected expleto got %s", cfg.AppName)
	}
}

func TestConfigEnvEmpty(t *testing.T) {

	os.Clearenv()
	cfg := DefaultConfig()
	if err := cfg.Sync(); err != nil {
		t.Errorf("Can't syncing env %v", err)
	}

	if cfg.Port != 9000 {
		t.Errorf("expected 9000 got %d instead", cfg.Port)
	}
}

func TestConfigEnvWrong(t *testing.T) {
	fields := []struct {
		name, env, value, error_msg string
	}{
		{"Port", "PORT", "--- 9009", fmt.Sprintf("expleto: loading config field %s %v", "Port", "strconv.ParseInt: parsing \"--- 9009\": invalid syntax")},
		{"Verbose", "VERBOSE", "true2", fmt.Sprintf("expleto: loading config field %s %v", "Verbose", "strconv.ParseBool: parsing \"true2\": invalid syntax")},
	}
	for _, f := range fields {
		os.Clearenv()

		cm := getEnvName(f.name)
		if cm != f.env {
			t.Errorf("expected %s got %s", f.env, cm)
		}

		os.Setenv(f.env, f.value)
		cfg := DefaultConfig()
		if err := cfg.Sync(); err.Error() != f.error_msg {

			t.Errorf("Got %v but expected %v", err, f.error_msg)
		}
	}
}

func TestGetEnvName(t *testing.T) {
	fixtures := []struct {
		name, env string
	}{
		{"AppName", "APP_NAME"},
		{"BaseURL", "BASE_URL"},
		{"Port", "PORT"},
		{"ViewsDir", "VIEWS_DIR"},
		{"StaticDir", "STATIC_DIR"},
		{"", ""},
	}
	for _, tt := range fixtures {
		result := getEnvName(tt.name)
		if result != tt.env {
			t.Fatal("Expected " + tt.env + " but got " + result)

		}
	}

}
