package expleto

//

import (
	"encoding/json"
	"errors"
	"fmt"
	// "io/ioutil"
	// "os"
	"path/filepath"
	// "reflect"
	// "strconv"
	// "strings"
	// "fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

// Config stores configurations values
type Config struct {
	AppName   string `json:"app_name" yaml:"app_name" toml:"app_name"`
	BaseURL   string `json:"base_url" yaml:"base_url" toml:"base_url"`
	Port      int    `json:"port" yaml:"port" toml:"port"`
	Verbose   bool   `json:"verbose" yaml:"verbose" toml:"verbose"`
	StaticDir string `json:"static_dir" yaml:"static_dir" toml:"static_dir"`
	ViewsDir  string `json:"view_dir" yaml:"view_dir" toml:"view_dir"`
}

// DefaultConfig returns the default configuration settings.
func DefaultConfig() *Config {
	return &Config{
		AppName:   "expleto web app",
		BaseURL:   "http://localhost:9000",
		Port:      9000,
		Verbose:   false,
		StaticDir: "static",
		ViewsDir:  "views",
	}
}

// NewConfig reads configuration from path. The format is deduced from the file extension
//	* .json    - is decoded as json
//	* .yml     - is decoded as yaml
//	* .toml    - is decoded as toml
func NewConfig(path string) (*Config, error) {
	var err error
	data, err := GetDataFromFile(path)
	if err != nil {
		return nil, FormatError(err)
	}

	cfg := &Config{}
	// err = nil
	switch filepath.Ext(path) {

	case ".json":
		err = json.Unmarshal(data, cfg)

	case ".toml":
		_, err = toml.Decode(string(data), cfg)

	case ".yml":
		err = yaml.Unmarshal(data, cfg)

	default:
		err = errors.New("expleto: The format is not supported")
	}
	if err != nil {
		return nil, fmt.Errorf("Can't parse %s: %v", path, err)
	}

	return cfg, nil
}
