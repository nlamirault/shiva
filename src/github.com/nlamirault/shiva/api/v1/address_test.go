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

package v1

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"

	"github.com/nlamirault/shiva/storage"
	"github.com/nlamirault/shiva/utils"
)

func checkResponseHeader(t *testing.T, header http.Header) {
	if header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf("Invalid content type: %v", header)
	}
}

func Test_GetUnknownMACAddress(t *testing.T) {
	db, err := storage.New("boltdb", &storage.Config{
		BackendURL: utils.Tempfile()})
	if err != nil {
		t.Fatalf("Can't create BoltDB backend.")
	}

	req, _ := http.NewRequest(echo.GET, "/01:23:45:67:89:ab", nil)
	rec := httptest.NewRecorder()

	ws := NewWebService(db)
	e := echo.New()
	e.Get("/:mac", ws.GetMacAddress)
	c := echo.NewContext(req, echo.NewResponse(rec, e), e)

	err = ws.GetMacAddress(c)
	if err != nil {
		t.Fatalf("API error : %v", err)
	}
	//resp := c.Response()

	checkResponseHeader(t, rec.HeaderMap)
	body := rec.Body.String()
	if body != "{\"error\":\"Unknown MAC \"}" ||
		rec.Code != http.StatusNotFound {
		t.Fatalf("Invalid API response: %s", body)
	}
}

func Test_GetValidMACAddress(t *testing.T) {
	db, err := storage.New("boltdb", &storage.Config{
		BackendURL: utils.Tempfile()})
	if err != nil {
		t.Fatalf("Can't create BoltDB backend.")
	}

	mac, err := net.ParseMAC("01:23:45:67:89:de")
	if err != nil {
		t.Fatalf("Can't parse MAC address : %v", err)
	}

	req, _ := http.NewRequest(echo.GET, fmt.Sprintf("/%s", mac.String()), nil)
	rec := httptest.NewRecorder()

	ws := NewWebService(db)
	e := echo.New()
	e.SetDebug(true)
	e.Get("/:mac", ws.GetMacAddress)
	c := echo.NewContext(req, echo.NewResponse(rec, e), e)

	data, err := json.Marshal(mac)
	if err != nil {
		t.Fatalf("Can't convert MAC to json : %v", err)
	}
	db.Put([]byte(mac.String()), data)

	err = ws.GetMacAddress(c)
	if err != nil {
		t.Fatalf("API error : %v", err)
	}
	//resp := c.Response()

	checkResponseHeader(t, rec.HeaderMap)
	body := rec.Body.String()
	if body != "{\n\"mac\": \"01:23:45:67:89:de\"\n}" ||
		rec.Code != http.StatusOK {
		t.Fatalf("Invalid API response: %s", body)
	}
}
