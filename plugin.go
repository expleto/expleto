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
	// 	"fmt"
	// 	"io/ioutil"
	// 	"log"
	// 	"net/rpc"
	// 	"os"
	// 	"os/exec"

	"github.com/hashicorp/go-plugin"
)

// PluginStore holds plugin names and their clients
type PluginStore struct {
	Plugins map[string]string
	Clients map[string]*plugin.Client
}
