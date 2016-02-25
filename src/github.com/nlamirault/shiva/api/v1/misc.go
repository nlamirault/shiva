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
	//"fmt"
	//"log"
	"net/http"

	"github.com/labstack/echo"
)

// Help send a message in JSON
func (ws *WebService) Help(c *echo.Context) error {
	return c.String(http.StatusOK,
		"Welcome to Shiva, a simple DHCP server\n")
}

// DisplayAPIVersion sends the API version in JSON format
func (ws *WebService) DisplayAPIVersion(c *echo.Context) error {
	return c.JSON(http.StatusOK, &APIVersion{Version: "1"})
}
