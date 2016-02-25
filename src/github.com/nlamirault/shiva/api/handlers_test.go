// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	//"fmt"
	"testing"

	"github.com/nlamirault/shiva/storage"
	"github.com/nlamirault/shiva/utils"
)

var api = map[string]string{
	"/":                    "GET",
	"/api/version":         "GET",
	"/api/v1/address":      "POST",
	"/api/v1/address/:mac": "GET",
}

func Test_WebServiceRoutes(t *testing.T) {
	db, err := storage.New("boltdb", &storage.Config{
		BackendURL: utils.Tempfile()})
	if err != nil {
		t.Fatalf("Can't create BoltDB backend.")
	}
	ws := GetWebService(db, nil)
	routes := ws.Routes()
	if len(routes) != 4 {
		t.Fatalf("Invalid routes. : %v", routes)
	}
	for _, route := range ws.Routes() {
		if api[route.Path] != route.Method {
			t.Fatalf("Unknown route. : %v", route)
		}
	}
}
