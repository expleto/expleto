package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	cfgFiles := []string{
		"fixtures/config/app.json",
		"fixtures/config/app.yml",
		"fixtures/config/app.toml",
	}

	cfg := DefaultConfig()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(dir)
	for _, f := range cfgFiles {
		nCfg, err := NewConfig(f)
		if err != nil {
			t.Fatal(err)
		}
		if nCfg.AppName != cfg.AppName {
			t.Errorf("expected %s got %s", cfg.AppName, nCfg.AppName)
		}
	}

}
