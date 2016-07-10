package expleto

import (
	// "os"
	// "path/filepath"
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
	cfgFiles := []string{
		"fixtures/config/app.json",
		"fixtures/config/app.yml",
		"fixtures/config/app.toml",
	}
	cfg := DefaultConfig()
	for _, f := range cfgFiles {
		nCfg, err := NewConfig(f)
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
}
