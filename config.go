package expleto

//

import (
	"encoding/json"
	"errors"
	"fmt"
	// "io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	// "fmt"
	"github.com/BurntSushi/toml"
	"github.com/fatih/camelcase"
	"gopkg.in/yaml.v2"
)

const (
	ERROR_FORMAT_NOT_SUPPORTED = "expleto: The format is not supported"
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
		err = errors.New(ERROR_FORMAT_NOT_SUPPORTED)
	}
	if err != nil {
		return nil, fmt.Errorf("Can't parse %s: %v", path, err)
	}

	return cfg, nil
}

// overrides c field's values that are set in the environment.
// The environment variable names are derived from config fields by underscoring, and uppercasing
// the name. E.g. AppName will have a corresponding environment variable APP_NAME
//
// NOTE only int, string and bool fields are supported and the corresponding values are set.
// when the field value is not supported it is ignored.
func (c *Config) Sync() error {
	cfg := reflect.ValueOf(c).Elem()
	cTyp := cfg.Type()

	for k := range make([]struct{}, cTyp.NumField()) {
		field := cTyp.Field(k)

		cm := getEnvName(field.Name)
		env := os.Getenv(cm)
		if env == "" {
			continue
		}
		switch field.Type.Kind() {
		case reflect.String:
			cfg.FieldByName(field.Name).SetString(env)
		case reflect.Int:
			v, err := strconv.Atoi(env)
			if err != nil {
				return fmt.Errorf("utron: loading config field %s %v", field.Name, err)
			}
			cfg.FieldByName(field.Name).Set(reflect.ValueOf(v))
		case reflect.Bool:
			b, err := strconv.ParseBool(env)
			if err != nil {
				return fmt.Errorf("utron: loading config field %s %v", field.Name, err)
			}
			cfg.FieldByName(field.Name).SetBool(b)
		}

	}
	return nil
}

// returns all upper case and underscore separated string, from field.
// field is a camel case string.
//
// example
//	AppName will change to APP_NAME
func getEnvName(field string) string {
	camSplit := camelcase.Split(field)
	var rst string
	for k, v := range camSplit {
		if k == 0 {
			rst = strings.ToUpper(v)
			continue
		}
		rst = rst + "_" + strings.ToUpper(v)
	}
	return rst
}
