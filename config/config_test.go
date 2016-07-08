package config

import (
	// "os"
	// "path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	cfgFiles := []string{
		"fixtures/config/app.json",
		"fixtures/config/app.yml",
		"fixtures/config/app.toml",
	}
	cfg := DefaultConfig()
	for _, f := range cfgFiles {
		t.Logf(f)
		nCfg, err := NewConfig(f)
		if err != nil {
			t.Fatal(err)
		}
		if nCfg.AppName != cfg.AppName {
			t.Errorf("expecetd %s got %s", cfg.AppName, nCfg.AppName)
		}
	}

}
