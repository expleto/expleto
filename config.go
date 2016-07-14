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
			// https://golang.org/pkg/reflect/#Value.SetString
			// SetString sets v's underlying value to x. It panics if v's Kind is not String or if CanSet() is false.
			cfg.FieldByName(field.Name).SetString(env)
		case reflect.Int:
			v, err := strconv.Atoi(env)
			if err != nil {
				return fmt.Errorf("expleto: loading config field %s %v", field.Name, err)
			}
			// https://golang.org/pkg/reflect/#Value.Set
			// Set assigns x to the value v. It panics if CanSet returns false. As in Go, x's value must be assignable to v's type.
			cfg.FieldByName(field.Name).Set(reflect.ValueOf(v))
		case reflect.Bool:
			b, err := strconv.ParseBool(env)
			if err != nil {
				return fmt.Errorf("expleto: loading config field %s %v", field.Name, err)
			}
			// https://golang.org/pkg/reflect/#Value.Bool
			// Bool returns v's underlying value. It panics if v's kind is not Bool.
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
