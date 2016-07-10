package expleto

import (
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
}
